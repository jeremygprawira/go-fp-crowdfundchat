package repository

import (
	"go-fp-crowdfundchat/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateDataToDB(user *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateDataToDB(user *model.User) (*model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}