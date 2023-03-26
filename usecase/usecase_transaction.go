package usecase

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
)

type TransactionUsecase interface{
	GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error)
}

type transactionUsecase struct {
	repo repository.BaseRepository
}

func NewTransactionUsecase(repo repository.BaseRepository) *transactionUsecase {
	return &transactionUsecase{repo}
}

func (u *transactionUsecase) GetTransactionByProjectID(request *model.ProjectTransactionListRequest) ([]*model.Transaction, error) {
	transaction, err := u.repo.FindTransactionByProjectID(request.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}