package handler

import (
	"go-fp-crowdfundchat/usecase/user"
	"go-fp-crowdfundchat/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
}