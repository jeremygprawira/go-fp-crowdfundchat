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

func (h *UserHandler) LoginUser(c *gin.Context) {
	var request model.LoginUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42202",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	successLoginUser, err := h.userUsecase.PostLoginUser(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40002",
			"responseMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "User has been logged in successfully",
		"name": successLoginUser.Name,
		"phone_no": successLoginUser.PhoneNo,
		"token": "noTokenExisted",
	})
}

func (h *UserHandler) IsPhoneNoAvailable(c *gin.Context) {
	var request model.PhoneNoBodyRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42203",
			"responseMessage": "Failed checking on phone number availability.",
		})
		return
	}

	verifyPhoneNumber, err := h.userUsecase.PostPhoneNoAvailability(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseCode": "50001",
			"responseMessage": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"isPhoneNoAvailable": verifyPhoneNumber,
	})
}

func (h *UserHandler) PinValidation(c *gin.Context) {
	var request model.PinValidationRequest
	
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42204",
			"responseMessage": "Failed checking on phone number availability.",
		})
		return
	}

	validatedPin, err := h.userUsecase.PostPinValidation(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40003",
			"responseMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"isPhoneNoAvailable": validatedPin,
	})
}

func (h *UserHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40004",
			"responseMessage": err.Error(),
		})
		return
	}

	path := "mock/images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40005",
			"responseMessage": err.Error(),
		})
		return
	}

	// JWT
	userID := 1

	_, err = h.userUsecase.PostUploadImage(userID, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40006",
			"responseMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"isPhoneNoAvailable": true,
	})
}