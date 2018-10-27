package middleware

import (
	"database/sql"
	"github.com/daveearley/product/app/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//DbTransaction wraps any non get request in a DB transaction
func DbTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := db.Begin()

		if c.Request.Method == http.MethodGet {
			return
		}

		if err != nil {
			response.Error(c, http.StatusInternalServerError, err)
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
