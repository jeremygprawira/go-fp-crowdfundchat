package usecase

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
)

type ProjectService interface{
	PostProjectList(userID int) ([]*model.Project, error)
}

type projectUsecase struct {
	repo repository.UserRepository
}

func NewProjectUsecase(repo repository.UserRepository) *projectUsecase {
	return &projectUsecase{repo}
}

func (u *projectUsecase) PostProjectList(userID int) ([]*model.Project, error) {
	if userID != 0 {
		projects, err := u.repo.FindProjectByUserID(userID)
		if err != nil {
			return projects, err
		}

		return projects, nil
	}

	projects, err := u.repo.FindAllProject()
	if err != nil {
		return projects, err
	}

	return projects, nil
}