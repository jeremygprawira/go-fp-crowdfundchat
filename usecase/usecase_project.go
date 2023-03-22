package usecase

import (
	"fmt"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"

	"github.com/gosimple/slug"
)

type ProjectUsecase interface{
	GetProjectList(userID int) ([]*model.ProjectListResponse, error)
	GetProjectDetail(request *model.ProjectDetailRequest) (*model.Project, error)
	PostCreateProject(request *model.CreateProjectRequest) (*model.Project, error)
	PutUpdateProject(requestID *model.UpdateProjectRequest, requestData *model.CreateProjectRequest) (*model.Project, error)
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

func (u *projectUsecase) GetProjectDetail(request *model.ProjectDetailRequest) (*model.Project, error) {
	project, err := u.repo.FindProjectByProjectID(request.ID)
	if err != nil {
		return project, err
	}

	return project, nil
}

func (u *projectUsecase) PostCreateProject(request *model.CreateProjectRequest) (*model.Project, error) {
	project := model.Project{}
	project.ProjectTitle = request.ProjectTitle
	project.ShortDescription = request.ShortDescription
	project.LongDescription = request.LongDescription
	project.Perks = request.Perks
	project.GoalAmount = request.GoalAmount
	project.UserID = request.User.ID
	
	projectSlugIdentifier := fmt.Sprintf("%s %d", request.ProjectTitle, request.User.ID)
	project.Slug = slug.Make(projectSlugIdentifier)

	newProject, err := u.repo.CreateProjectToDB(&project)
	if err != nil {
		return newProject, err
	}

	return newProject, nil
}

func (u *projectUsecase) PutUpdateProject(requestID *model.UpdateProjectRequest, requestData *model.CreateProjectRequest) (*model.Project, error) {
	request, err := u.repo.FindProjectByProjectID(requestID.ID)
	if err != nil {
		return request, nil
	}

	if request.UserID != requestData.User.ID {
		return request, fmt.Errorf("%d, is not the owner of this campaign", request.UserID)
	}

	request.ProjectTitle = requestData.ProjectTitle
	request.ShortDescription = requestData.ShortDescription
	request.LongDescription = requestData.LongDescription
	request.Perks = requestData.Perks
	request.GoalAmount = requestData.GoalAmount

	requestedUpdate, err := u.repo.UpdateProjectToDB(request)
	if err != nil {
		return requestedUpdate, err
	}

	return requestedUpdate, nil
}