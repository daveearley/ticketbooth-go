package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, v := range c.Errors {
			fmt.Println(v.Err.Error())
		}

		return
	}
}
