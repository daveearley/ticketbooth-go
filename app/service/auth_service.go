package service

import (
	"../../app"
	"../../app/api/request"
	"../../app/repository"
	"../../app/utils"
	"../../configs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthService interface {
	ValidateLoginAndReturnJwtToken(req *request.Login) (string, error)
}

type userService struct {
	ur     repository.UserRepository
	config *configs.Config
}

func NewAuthService(r repository.UserRepository, c *configs.Config) AuthService {
	return &userService{r, c}
}

// ValidateLoginAndReturnJwtToken accepts returns a JWT token when a valid email/password combo is passed
func (s *userService) ValidateLoginAndReturnJwtToken(req *request.Login) (string, error) {
	u, err := s.ur.FindByEmail(req.Email)

	if err != nil {
		return "", err
	}

	if utils.CheckPasswordHash(req.Password, u.Password) == false {
		return "", app.UnauthorizedError("JWT Login")
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

	return token.SignedString(mySigningKey)
}
