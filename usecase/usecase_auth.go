package usecase

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GenerateToken(userID int) (string, error)
}

type authUsecase struct {

}

var AuthSecretKey = []byte(os.Getenv("SECRET_KEY"))
// var AuthSecretKey = "goG0_fpFP_cR0WDfUnd1N6"


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