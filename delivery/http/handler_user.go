package http

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.Usecase
}

func NewUserHandler (userUsecase usecase.Usecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userRequest model.RegisterUserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42201",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	user, err := h.userUsecase.PostRegisterUser(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40001",
			"responseMessage": "Usecase PostRegisterUser is not working properly",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "User has been registered successfully",
		"name": user.Name,
		"phone_no": user.PhoneNo,
		"pin": user.PIN,
		"token": "noTokenExisted",
	})
}






/*
type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser (c *gin.Context) {
	var userRequest user.RegisterUserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {

		rawFailedMessage := util.FailedResponse(err) 
		failedMessage := gin.H{"errors": rawFailedMessage}
		userResponse := util.JSON(http.StatusBadRequest, "failed", "Account has not been registered", failedMessage)
		c.JSON(http.StatusBadRequest, userResponse)
		return
	}

	registeredUser, err := h.userService.PostRegisterUser(userRequest)
	if err != nil {
		userResponse := util.JSON(http.StatusBadRequest, "failed", "Account has not been registered", nil)
		c.JSON(http.StatusBadRequest, userResponse)
		return
	}
	
	rawResponse := user.UserResponse(registeredUser, "jwtToken")
	userResponse := util.JSON(http.StatusOK, "success", "Account has been registered", rawResponse)

	c.JSON(http.StatusOK, userResponse)
}*/