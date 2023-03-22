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
	PostUploadProjectImage(request *model.UploadProjectImageRequest, fileLocation string) (*model.ProjectImages, error)
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
		return request, fmt.Errorf("this user ID: %d - is not the owner of this campaign", request.UserID)
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

func (u *projectUsecase) PostUploadProjectImage(request *model.UploadProjectImageRequest, fileLocation string) (*model.ProjectImages, error) {
	project, err := u.repo.FindProjectByProjectID(request.ProjectID)
	if err != nil {
		return &model.ProjectImages{}, err
	}

	if project.UserID != request.User.ID {
		return &model.ProjectImages{}, fmt.Errorf("this user ID: %d - is not the owner of this campaign", request.User.ID)
	}

	isPrimary := 0
	if request.IsPrimary {
		isPrimary = 1
		_, err := u.repo.SetImageToNonPrimary(request.ProjectID)
		if err != nil {
			return &model.ProjectImages{}, err
		}
	}

	projectImage := model.ProjectImages{}
	projectImage.ProjectID = request.ProjectID
	projectImage.IsPrimary = isPrimary
	projectImage.FileName = fileLocation

	newProjectImage, err := u.repo.UploadImageToDB(&projectImage)
	if err != nil {
		return newProjectImage, err
	}

	return newProjectImage, nil
}