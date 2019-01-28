package middleware

import (
	"database/sql"
	"../../../app"
	"../../../app/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//DbTransaction wraps any non GET request in a DB transaction. If any error is detected the transaction
// will roll back
func DbTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := db.Begin()

		if c.Request.Method == http.MethodGet {
			return
		}

		if err != nil {
			response.Error(c, app.ServerError(err))
			return
		}

		c.Next()

		if len(c.Errors) > 0 {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
