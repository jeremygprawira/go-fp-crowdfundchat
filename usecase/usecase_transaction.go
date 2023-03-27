package usecase

import (
	"errors"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
)

type TransactionUsecase interface{
	GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error)
	GetTransactionByUserID(userID int) ([]*model.Transaction, error)
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