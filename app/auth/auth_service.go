package auth

import (
	"errors"
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/users"
	"github.com/daveearley/product/app/utils"
	"github.com/daveearley/product/configs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service interface {
	ValidateLoginAndReturnJwtToken(req *request.Login) (string, error)
}

type service struct {
	ur     users.Repository
	config *configs.Config
}

func NewService(r users.Repository, c *configs.Config) Service {
	return &service{r, c}
}

// ValidateLoginAndReturnJwtToken accepts returns a JWT token when a valid email/password combo is passed
func (s *service) ValidateLoginAndReturnJwtToken(req *request.Login) (string, error) {
	u, err := s.ur.FindByEmail(req.Username)

	if err != nil {
		return "", err
	}

	if utils.CheckPasswordHash(req.Password, u.Password) == false {
		return "", errors.New("incorrect password")
	}

	mySigningKey := []byte(s.config.JwtSecret)

	type jwtClaims struct {
		UserId    int `json:"user_id"`
		AccountId int `json:"account_id"`
		jwt.StandardClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		u.ID,
		u.AccountID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    "api",
		},
	})

	ss, err := token.SignedString(mySigningKey)

	return ss, err
}
