package app

import (
	"fmt"
	"io"
	"net/http"
)

const (
	TypeUnauthorized   = "UnauthorizedError"
	TypeNotFound       = "NotFoundError"
	TypeServerError    = "ServerError"
	TypeMissingField   = "MissingFieldError"
	TypeInvalidValue   = "InvalidValueError"
	TypeInvalidRequest = "InvalidRequest"
)

// Error provides some additional on top of Go's generic error
type Error struct {
	Code     int                    `json:"-"`
	Type     string                 `json:"type,omitempty"`
	Context  map[string]interface{} `json:"context,omitempty"`
	Message  string                 `json:"message,omitempty"`
	original error
}

// Error shows a string representation of the error
func (err Error) Error() string {
	for key, val := range err.Context {
		fmt.Printf("%s='%s' ", key, val)
	}
	return fmt.Sprintf("%s: %s", err.Type, err.Message)
}

//Format implements the fmt Formatter interface
func (err Error) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, err.Error())
			return
		} else {
			fmt.Fprintf(st, "%s: ", err.Type)
			for key, val := range err.Context {
				fmt.Fprintf(st, "%s='%s' ", key, val)
			}
			return
		}
	case 's':
		io.WriteString(st, err.Error())
	case 'q':
		fmt.Fprintf(st, "%q", err.Error())
	}
}

// ServerError returns a server error with additional args
func ServerError(err error, args ...interface{}) error {
	return &Error{
		Code: http.StatusInternalServerError,
		Type: TypeServerError,
		Context: map[string]interface{}{
			"error": err,
			"args":  fmt.Sprint(args...),
		},
		Message: "Server Error",
	}
}

// UnauthorizedError returns an unauthorized error
func UnauthorizedError(action string) error {
	return &Error{
		Code: http.StatusUnauthorized,
		Type: TypeUnauthorized,
		Context: map[string]interface{}{
			"action": action,
		},
		Message: "Action is unauthorized",
	}
}

// NotFoundError returns a not found error
func NotFoundError(resource string, resourceID interface{}) error {
	return &Error{
		Code: http.StatusNotFound,
		Type: TypeNotFound,
		Context: map[string]interface{}{
			"resource":    resource,
			"resource_id": resourceID,
		},
		Message: "Unable to find resource",
	}
}

// InvalidValueError can be used to generate an error that represents an invalid
// value for the 'field'. reason should be used to add detail describing why
// the value is invalid.
func InvalidValueError(field string, reason string) error {
	return &Error{
		Code:    http.StatusBadRequest,
		Type:    TypeInvalidValue,
		Message: "A parameter has invalid value",
		Context: map[string]interface{}{
			"field":  field,
			"reason": reason,
		},
	}
}

// MissingFieldError can be used to generate an error that represents
// a empty value for a required field.
func MissingFieldError(field string) error {
	return &Error{
		Code:    http.StatusBadRequest,
		Type:    TypeMissingField,
		Message: "A required field is missing",
		Context: map[string]interface{}{
			"field": field,
		},
	}
}

// ValidationError returns an error that can be used to represent an invalid request.
func ValidationError(reason string) error {
	return &Error{
		Code:    http.StatusBadRequest,
		Type:    TypeInvalidRequest,
		Message: reason,
		Context: map[string]interface{}{},
	}
}
