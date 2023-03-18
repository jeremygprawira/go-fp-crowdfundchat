package usecase

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
)

type ProjectUsecase interface{
	GetProjectList(userID int) ([]*model.ProjectListResponse, error)
}

type projectUsecase struct {
	repo repository.BaseRepository
}

func NewProjectUsecase(repo repository.BaseRepository) *projectUsecase {
	return &projectUsecase{repo}
}

func (u *projectUsecase) GetProjectList(userID int) ([]*model.ProjectListResponse, error) {
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