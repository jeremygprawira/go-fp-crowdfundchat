package usecase

import (
	"errors"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
	"strconv"
)

type TransactionUsecase interface{
    GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error)
    GetTransactionByUserID(userID int) ([]*model.Transaction, error)
    PostOrderTransaction(request *model.OrderTransactionRequest) (*model.Transaction, error)
	PostVerifyTransaction(request *model.TransactionNotificationRequest) error
}

type transactionUsecase struct {
    repo repository.BaseRepository
}

func NewTransactionUsecase(repo repository.BaseRepository) *transactionUsecase {
    return &transactionUsecase{repo}
}

func (u *transactionUsecase) GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error) {
    project, err := u.repo.FindProjectByProjectID(request.ID)
    if err != nil {
        return []*model.Transaction{}, err
    }

    if project.UserID != request.User.ID {
        return []*model.Transaction{}, errors.New("not the owner of the project")
    }
    
    transaction, err := u.repo.FindTransactionByProjectID(request.ID)
    if err != nil {
        return transaction, err
    }

    return transaction, nil
}

func (u *transactionUsecase) GetTransactionByUserID(userID int) ([]*model.Transaction, error) {
    transaction, err := u.repo.FindTransactionByUserID(userID)
    if err != nil {
        return transaction, err
    }

    return transaction, nil
}

func (u *transactionUsecase) PostOrderTransaction(request *model.OrderTransactionRequest) (*model.Transaction, error) {
    transaction := model.Transaction{}
    transaction.ProjectID = request.ProjectID
    transaction.Amount = request.Amount
    transaction.UserID = request.User.ID
    transaction.Status = "PENDING"
    transaction.ReceiptNo = ""

    orderedTransaction, err := u.repo.CreateTransactionToDB(&transaction)
    if err != nil {
        return orderedTransaction, err
    }

    transactionURL, err := util.GetTransactionURL(orderedTransaction, &request.User)
    if err != nil {
        return orderedTransaction, err
    }

    orderedTransaction.TransactionURL = transactionURL
    
    orderedTransaction, err = u.repo.UpdateTransactionToDB(orderedTransaction)
    if err != nil {
        return orderedTransaction, err
    }

    return orderedTransaction, nil
}

func (u *transactionUsecase) PostVerifyTransaction(request *model.TransactionNotificationRequest) error {
	transaction_id, _ := strconv.Atoi(request.OrderID)

	transaction, err := u.repo.FindTransactionByID(transaction_id)
	if err != nil {
		return err
	}

	if (request.PaymentType == "credit_card" && request.TransactionStatus == "capture" && request.FraudStatus == "accept") {
		transaction.Status = "SUCCESS"
	} else if (request.TransactionStatus == "settlement") {
		transaction.Status = "SUCCESS"
	} else if (request.TransactionStatus == "deny" || request.TransactionStatus == "expire" || request.TransactionStatus == "cancel") {
		transaction.Status = "CANCELLED"
	}

	verifiedTransaction, err := u.repo.UpdateTransactionToDB(transaction)
	if err != nil {
		return err
	}

	project, err := u.repo.FindProjectByProjectID(verifiedTransaction.ProjectID)
	if err != nil {
		return err
	}

	if verifiedTransaction.Status == "SUCCESS" {
		project.ContributorCount += 1
		project.CurrentAmount += verifiedTransaction.Amount

		_, err := u.repo.UpdateProjectToDB(project)
		if err != nil {
			return err
		}
	}

	return nil
}