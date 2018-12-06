package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getID(p *gin.Param) int {
	id, _ := strconv.Atoi(p.Value)

	return id
}
