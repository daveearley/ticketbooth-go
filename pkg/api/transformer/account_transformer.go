package transformer

import (
	"github.com/daveearley/product/pkg/model"
	"github.com/gin-gonic/gin"
)

func AccountTransformer(ac *model.Account) *gin.H {
	return &gin.H{
		"data": ac,
	}
}
