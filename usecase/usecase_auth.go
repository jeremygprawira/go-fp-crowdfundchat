package usecase

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authUsecase struct {

}

var AuthSecretKey = []byte(os.Getenv("SECRET_KEY"))


func NewAuthUsecase() *authUsecase {
	return &authUsecase{}
}

func (s *authUsecase) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	signedToken, err := token.SignedString(AuthSecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authUsecase) ValidateToken(encodedToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
	
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(AuthSecretKey), nil
	})

	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}