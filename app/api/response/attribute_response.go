package response

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
	for _, attr := range attrs {
		transformed = append(transformed, &AttributeResponse{
			ID:    attr.ID,
			Name:  attr.Name,
			Value: attr.Value,
		})
	}

	return &transformed
}
