package transformer

import (
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/gin-gonic/gin"
)

func AccountTransformer(ac *models.Account) *gin.H {
	return &gin.H{
		"data": ac,
	}
}
