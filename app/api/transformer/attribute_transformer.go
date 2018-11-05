package transformer

import (
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

type AttributeResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TransformAttributes(c *gin.Context, attrs []*models.Attribute) interface{} {
	var transformed []interface{}
	for _, v := range attrs {
		transformed = append(transformed, &AttributeResponse{
			ID:    v.ID,
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return &transformed
}
