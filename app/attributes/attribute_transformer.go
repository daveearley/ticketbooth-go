package attributes

import (
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TransformMany(c *gin.Context, attrs []*models.Attribute) []*Response {
	var transformed []*Response
	for _, v := range attrs {
		transformed = append(transformed, &Response{
			ID:    v.ID,
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return transformed
}
