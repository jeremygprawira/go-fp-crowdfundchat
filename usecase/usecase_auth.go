package usecase

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetToken(c *gin.Context) string
	//GetTokenID(c *gin.Context) (int, error)
	GetUserIDByToken(parsedToken *jwt.Token) (int, error)
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

func (u *authUsecase) GetUserIDByToken(parsedToken *jwt.Token) (int, error) {

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok && !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	return int(claims["user_id"].(float64)), nil
}
