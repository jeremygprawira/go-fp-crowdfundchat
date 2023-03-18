package usecase

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetToken(c *gin.Context) string
	GetTokenID(c *gin.Context) (int, error)
}

type authUsecase struct {
	secretKey	string
}

func NewAuthUsecase() *authUsecase {
	return &authUsecase{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (u *authUsecase) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (u *authUsecase) ValidateToken(encodedToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
	
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte([]byte(os.Getenv("SECRET_KEY"))), nil
	})

	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}

func (u *authUsecase) GetToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (u *authUsecase) GetTokenID(c *gin.Context) (int, error) {

	tokenString := u.GetToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(uid), nil
	}
	return 0, nil
}