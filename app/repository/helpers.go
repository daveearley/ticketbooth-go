package repository

import (
	"database/sql"
	"../../app"
	"github.com/pkg/errors"
)

// getErrorType accepts an error from SQLBoiler and checks if it's server error or not rows found error
func getErrorType(err error, resourceType string, resourceID interface{}) error {
	if errors.Cause(err) == sql.ErrNoRows {
		return app.NotFoundError(resourceType, resourceID)
	}

	return app.ServerError(err, resourceType, resourceID)
}
