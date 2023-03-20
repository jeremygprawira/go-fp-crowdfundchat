package middleware

import (
	"go-fp-crowdfundchat/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := usecase.NewAuthUsecase().GetToken(c)
		receivedToken, err := usecase.NewAuthUsecase().ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"responseCode": "40101",
				"responseMessage": "Unauthorized",
			})
			return
		}

		userID, err := usecase.NewAuthUsecase().GetUserIDByToken(receivedToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"responseCode": "40102",
				"responseMessage": "Unauthorized",
			})
			return
		}

		c.Set("currentUser", userID)
		c.Next()
		return
	}

}