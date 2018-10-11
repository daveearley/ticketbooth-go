package service

import (
	"errors"
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/repository"
	"github.com/daveearley/product/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type authService struct {
	ur repository.UserRepository
}

func NewAuthService(r repository.UserRepository) *authService {
	return &authService{r}
}

// ValidateLoginAndReturnJwtToken accepts returns a JWT token when a valid email/password combo is passed
func (s *authService) ValidateLoginAndReturnJwtToken(req *request.Login) (string, error) {
	user, err := s.ur.FindByEmail(req.Username)

	if err != nil {
		return "", err
	}

	if utils.CheckPasswordHash(req.Password, user.Password) == false {
		return "", errors.New("incorrect password")
	}

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	type jwtClaims struct {
		UserId    int `json:"user_id"`
		AccountId int `json:"account_id"`
		jwt.StandardClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		user.ID,
		user.AccountID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    "api",
		},
	})

	ss, err := token.SignedString(mySigningKey)

	return ss, err
}
