package middleware

import (
	"go-fp-crowdfundchat/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := usecase.NewAuthUsecase().GetToken(c)
		receivedToken, err := usecase.NewAuthUsecase().ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"responseCode": "40107",
				"responseMessage": "Unauthorized",
			})
			return
		}
		_, ok := receivedToken.Claims.(jwt.MapClaims)
		if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"responseCode": "40108",
			"responseMessage": "Unauthorized",
		})
		return
	}

	}
}