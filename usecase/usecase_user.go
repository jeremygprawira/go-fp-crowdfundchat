package usecase

import (
	"errors"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
)

type UserUsecase interface {
	PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error)
	PostLoginUser(request *model.LoginUserRequest) (*model.User, error)
	PostPhoneNoAvailability(request *model.PhoneNoBodyRequest) (bool, error)
	PostPinValidation(request *model.PinValidationRequest) (bool, error)
	PostUploadImage(userID int, fileLocation string) (*model.User, error)
 	GetUserByID(userID int) (*model.User, error)
	GetOnlineUserByStatus(status string) (*model.User, error)
	PostUpdateUserStatus(userID int, status string) (*model.User, error)
}

type userUsecase struct {
	repo repository.BaseRepository
}

func NewUserUsecase(repo repository.BaseRepository) *userUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error) {
	user := model.User{}
	user.Name = request.Name
	user.PhoneNo = request.PhoneNo
	encrpytedPassword, err := util.HashSaltPassword(request.PIN)
	if err != nil {
		return nil, err
	}
	user.PIN = string(encrpytedPassword)
	user.Role = "user"
	user.Status = "offline"

	newUser, err := u.repo.CreateDataToDB(&user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *userUsecase) PostLoginUser(request *model.LoginUserRequest) (*model.User, error) {
	
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

	user.Status = "online"
	updatedStatus, err := u.repo.UpdateDataToDB(user)
	if err != nil {
		return user, err
	}

	return updatedStatus, nil
}

func (u *userUsecase) PostPhoneNoAvailability(request *model.PhoneNoBodyRequest) (bool, error) {
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

func (u *userUsecase) PostPinValidation(request *model.PinValidationRequest) (bool, error) {
	pin := request.PIN
	_, err := util.RepeatingPINChecker(pin)
	if err != nil {
		return false, errors.New("PIN cannot be 111111")
	}

	return true, nil
}

func (u *userUsecase) PostUploadImage(userID int, fileLocation string) (*model.User, error) {
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

func (u *userUsecase) GetUserByID(userID int) (*model.User, error) {
	user, err := u.repo.FindUserByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("userID is not found")
	}

	return user, nil
}

func(u *userUsecase) GetOnlineUserByStatus(status string) (*model.User, error) {
	user, err := u.repo.FindStatus(status)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) PostUpdateUserStatus(userID int, status string) (*model.User, error) {
	user, err := u.repo.FindUserByID(userID)
	if err != nil {
		return user, errors.New("cannot find the user ID")
	}

	user.Status = status
	updatedStatus, err := u.repo.UpdateDataToDB(user)
	if err != nil {
		return updatedStatus, err
	}

	return updatedStatus, nil
}