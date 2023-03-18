package repository

import (
	"go-fp-crowdfundchat/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateDataToDB(user *model.User) (*model.User, error)
	FindDataByPhoneNo(phoneNo string) (*model.User, error)
	FindUserByID(Id int) (*model.User, error)
	UpdateDataToDB(user *model.User) (*model.User, error)

	FindAllProject() ([]*model.Project, error)
	FindProjectByUserID(userID int) ([]*model.Project, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateDataToDB(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindDataByPhoneNo(phoneNo string) (*model.User, error) {
	var user *model.User
	err := r.db.Where("phone_no = ?", phoneNo).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByID(Id int) (*model.User, error) {
	var user *model.User
	err := r.db.Where("id = ?", Id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateDataToDB(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindAllProject() ([]*model.Project, error) {
	var projects []*model.Project
	err := r.db.Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (r *userRepository) FindProjectByUserID(userID int) ([]*model.Project, error) {
	var projects []*model.Project
	
	err := r.db.Where("user_id = ?", userID).Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}