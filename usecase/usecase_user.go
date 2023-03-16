package usecase

import (
	"errors"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
)

type Usecase interface {
	PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error)
	PostLoginUser(request *model.LoginUserRequest) (*model.User, error)
	PostPhoneNoAvailability(request *model.PhoneNoBodyRequest) (bool, error)
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

func (u *usecase) PostLoginUser(request *model.LoginUserRequest) (*model.User, error) {
	
	phoneNo := request.PhoneNo
	pin := request.PIN

	user, err := u.repo.FindDataByPhoneNo(phoneNo)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that phone number")
	}

	err = util.CheckPINCredibility(pin, user.PIN)
	if err != nil {
		return user, errors.New("pin is incorrect")
	}

	return user, nil
}

func (u *usecase) PostPhoneNoAvailability(request *model.PhoneNoBodyRequest) (bool, error) {
	phoneNo := request.PhoneNo
	user, err := u.repo.FindDataByPhoneNo(phoneNo)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}