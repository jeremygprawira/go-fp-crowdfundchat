package repository

import (
	"go-fp-crowdfundchat/model"

	"gorm.io/gorm"
)

type BaseRepository interface {
	CreateDataToDB(user *model.User) (*model.User, error)
	FindDataByPhoneNo(phoneNo string) (*model.User, error)
	FindUserByID(Id int) (*model.User, error)
	UpdateDataToDB(user *model.User) (*model.User, error)

	FindAllProject() ([]*model.ProjectListResponse, error)
	FindProjectByUserID(userID int) ([]*model.ProjectListResponse, error)
	FindProjectByProjectID(ID int) ([]*model.ProjectDetailResponse, error)
	CreateProjectToDB(project *model.Project) (*model.Project, error)
}

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *baseRepository {
	return &baseRepository{db}
}

func (r *baseRepository) CreateDataToDB(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *baseRepository) FindDataByPhoneNo(phoneNo string) (*model.User, error) {
	var user *model.User
	err := r.db.Where("phone_no = ?", phoneNo).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *baseRepository) FindUserByID(Id int) (*model.User, error) {
	var user *model.User
	err := r.db.Where("id = ?", Id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *baseRepository) UpdateDataToDB(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *baseRepository) FindAllProject() ([]*model.ProjectListResponse, error) {
	var projects []*model.Project
	var projectResponse []*model.ProjectListResponse
	//err := r.db.Table("projects").Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	err := r.db.Model(&projects).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projectResponse).Error
	if err != nil {
		return projectResponse, err
	}

	return projectResponse, nil
}

func (r *baseRepository) FindProjectByUserID(userID int) ([]*model.ProjectListResponse, error) {
	var projects []*model.Project
	var projectResponse []*model.ProjectListResponse
	//err := r.db.Table("projects").Where("user_id = ?", userID).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	err := r.db.Model(&projects).Where("user_id =?", userID).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projectResponse).Error
	if err != nil {
		return projectResponse, err
	}

	return projectResponse, nil
}

func (r *baseRepository) FindProjectByProjectID(ID int) ([]*model.ProjectDetailResponse, error) {
	var projects []*model.Project
	var projectResponse []*model.ProjectDetailResponse
	//err := r.db.Table("projects").Where("user_id = ?", userID).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	err := r.db.Model(&projects).Where("id =?", ID).Preload("User").Preload("ProjectImages").Find(&projectResponse).Error
	if err != nil {
		return projectResponse, err
	}

	return projectResponse, nil
}

func (r *baseRepository) CreateProjectToDB(project *model.Project) (*model.Project, error) {
	err := r.db.Save(&project).Error
	if err != nil {
		return project, err
	}

	return project, nil
}