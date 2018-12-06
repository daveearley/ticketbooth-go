package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"errors"
	"gopkg.in/go-playground/validator.v9"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

var (
	ErrorInternalError = errors.New("Woops! Something went wrong :(")
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			//list := make(map[string]string)

			//for _, e := range c.Errors {

			//errs := e.Err.(validator.ValidationErrors)
			//for _, err := range errs {
			//	list[err.Field()] = ValidationErrorToText(err)
			//}

			// Make sure we maintain the preset response status
			//	status := http.StatusBadRequest
			//	if c.Writer.Status() != http.StatusOK {
			//		status = c.Writer.Status()
			//	}
			//	c.JSON(status, gin.H{"Errors": list})
			//
			//}
			//// If there was no public or bind error, display default 500 message
			//if !c.Writer.Written() {
			//	c.JSON(http.StatusInternalServerError, gin.H{"error": ErrorInternalError.Error()})
			//}
			//fmt.Println(list)
		}
	}
}

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	case "gtefield":
		return fmt.Sprintf("%s must after %s", e.Field(), e.Param())
	case "ltefield":
		return fmt.Sprintf("%s must before %s", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}
