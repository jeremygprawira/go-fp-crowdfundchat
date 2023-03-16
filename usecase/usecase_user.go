package usecase

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
)

type Usecase interface {
	PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error)
}

type usecase struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *usecase {
	return &usecase{repo}
}

func (u *usecase) PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error) {
	user := model.User{}
	user.Name = request.Name
	user.PhoneNo = request.PhoneNo
	encrpytedPassword, err := util.HashSaltPassword(request.PIN)
	if err != nil {
		return nil, err
	}
	user.PIN = string(encrpytedPassword)
	user.Role = "user"

	newUser, err := u.repo.CreateDataToDB(&user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}










/*
type Service interface {
	PostRegisterUser (input RegisterUserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) PostRegisterUser (request RegisterUserRequest) (User, error) {
	user := User{}
	user.Name = request.Name
	user.PhoneNo = request.PhoneNo
	// user.PIN = util.HashPassword(request.PIN, util.GenerateRandomSalt(util.SaltSize))
	encrpytedPassword, err := util.HashSaltPassword(request.PIN)
	if err != nil {
		return user, err
	}
	user.PIN = string(encrpytedPassword)
	user.Role = "user"

	newUser, err := s.repository.RegisterToDB(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}*/