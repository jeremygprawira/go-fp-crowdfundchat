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

	FindAllProject() ([]*model.Project, error)
	FindProjectByUserID(userID int) ([]*model.Project, error)
	//FindProjectByProjectID(ID int) ([]*model.ProjectDetailResponse, error)
	FindProjectByProjectID(ID int) (*model.Project, error)
	CreateProjectToDB(project *model.Project) (*model.Project, error)
	UpdateProjectToDB(project *model.Project) (*model.Project, error)
	UploadImageToDB(projectImage *model.ProjectImages) (*model.ProjectImages, error)
	SetImageToNonPrimary(projectID int) (bool, error)

	FindTransactionByProjectID(projectID int) ([]*model.Transaction, error)
	FindTransactionByUserID(userID int) ([]*model.Transaction, error)
	CreateTransactionToDB(transaction *model.Transaction) (*model.Transaction, error)
	UpdateTransactionToDB(transaction *model.Transaction) (*model.Transaction, error)
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

func (r *baseRepository) FindAllProject() ([]*model.Project, error) {
	var projects []*model.Project
	//var projectResponse []*model.ProjectListResponse
	//err := r.db.Table("projects").Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	err := r.db.Model(&projects).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (r *baseRepository) FindProjectByUserID(userID int) ([]*model.Project, error) {
	var projects []*model.Project
	//var projectResponse []*model.ProjectListResponse
	err := r.db.Model(&projects).Table("projects").Where("user_id = ?", userID).Preload("ProjectImages", "project_images.is_primary = 1").Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (r *baseRepository) FindProjectByProjectID(ID int) (*model.Project, error) {
	var projects *model.Project
	err := r.db.Model(&projects).Where("id =?", ID).Preload("User").Preload("ProjectImages").Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (r *baseRepository) CreateProjectToDB(project *model.Project) (*model.Project, error) {
	err := r.db.Save(&project).Error
	if err != nil {
		return project, err
	}

	return project, nil
}

func (r *baseRepository) UpdateProjectToDB(project *model.Project) (*model.Project, error) {
	err := r.db.Save(&project).Error
	if err != nil {
		return project, err
	}

	return project, nil
}

func (r *baseRepository) UploadImageToDB(projectImage *model.ProjectImages) (*model.ProjectImages, error) {
	err := r.db.Save(&projectImage).Error
	if err != nil {
		return projectImage, err
	}

	return projectImage, nil
}

func (r *baseRepository) SetImageToNonPrimary(projectID int) (bool, error) {
	var projectImage *model.ProjectImages
	err := r.db.Model(&projectImage).Where("project_id = ?", projectID).Update("is_primary", 0).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *baseRepository) FindTransactionByProjectID(projectID int) ([]*model.Transaction, error) {
	var transaction []*model.Transaction
	err := r.db.Model(&transaction).Preload("User").Where("project_id = ?", projectID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *baseRepository) FindTransactionByUserID(userID int) ([]*model.Transaction, error) {
	var transaction []*model.Transaction
	err := r.db.Preload("Project.ProjectImages", "project_images.is_primary = 1").Where("user_id =?", userID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	
	return transaction, nil
}

func (r *baseRepository) CreateTransactionToDB(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *baseRepository) UpdateTransactionToDB(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}