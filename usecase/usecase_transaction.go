package usecase

import (
	"errors"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
)

type TransactionUsecase interface{
	GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error)
	GetTransactionByUserID(userID int) ([]*model.Transaction, error)
	PostOrderTransaction(request *model.OrderTransactionRequest) (*model.Transaction, error)
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