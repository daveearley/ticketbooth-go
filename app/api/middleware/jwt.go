package middleware

import (
	"errors"
	"github.com/daveearley/ticketbooth/app/repository"
	"github.com/daveearley/ticketbooth/configs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtMiddleware(repository repository.UserRepository, config *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := getTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			u, err := repository.GetById(int(claims["user_id"].(float64)))

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Authorization token belongs to non-existent user",
				})
			}

			c.Set("auth_user", u)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func getTokenFromHeader(token string) (string, error) {
	if len(token) > 6 && strings.ToLower(token[0:7]) == "bearer " {
		return token[7:], nil
	}
	return token, errors.New("No Authorization header in request")
}
