package service

import (
	"fmt"
	"github.com/daveearley/product/pkg/models/generated"
)

// This maps a map[string]interface{} to Attribute models
// todo Move attributes mapping to utils. Also handle conversion of map to json string
func MapToAttributes(attributes *map[string]interface{}) []*models.Attribute {
	var attrs []*models.Attribute
	for k, v := range *attributes {

		//if reflect.TypeOf(v).Kind().String() == "map" {
		//	v, _ := json.Marshal(v)
		//}

		attrs = append(attrs, &models.Attribute{
			Name:  k,
			Value: fmt.Sprint(v),
		})
	}

	return attrs
}
