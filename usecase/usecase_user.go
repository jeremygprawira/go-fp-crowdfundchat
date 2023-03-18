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
	PostPinValidation(request *model.PinValidationRequest) (bool, error)
	PostUploadImage(userID int, fileLocation string) (*model.User, error)
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

func (u *usecase) PostPinValidation(request *model.PinValidationRequest) (bool, error) {
	pin := request.PIN
	_, err := util.RepeatingPINChecker(pin)
	if err != nil {
		return false, errors.New("PIN cannot be 111111")
	}

	return true, nil
}

func (u *usecase) PostUploadImage(userID int, fileLocation string) (*model.User, error) {
	user, err := u.repo.FindUserByID(userID)
	if err != nil {
		return user, errors.New("cannot find the user ID")
	}

	user.AvatarFileName = fileLocation
	uploadedPhoto, err := u.repo.UpdateDataToDB(user)
	if err != nil {
		return uploadedPhoto, errors.New("failed to upload the photo to database")
	}

	return uploadedPhoto, nil
}