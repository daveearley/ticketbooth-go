package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		//todo handle errors
		//for _, v := range c.Errors {

		//}

		return
	}
}
