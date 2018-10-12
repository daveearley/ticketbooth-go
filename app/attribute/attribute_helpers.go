package attribute

import (
	"encoding/json"
	"fmt"
	"github.com/daveearley/product/app/models/generated"
	"reflect"
)

// This maps a map[string]interface{} to Attribute models
func MapToAttributes(attributes *map[string]interface{}) []*models.Attribute {
	var attrs []*models.Attribute
	for k, v := range *attributes {
		v, t := valueToStringAndType(v)
		attrs = append(attrs, &models.Attribute{
			Name:  k,
			Value: v,
			Type:  t,
		})
	}

	return attrs
}

// Casts value to string and returns the string + values type
func valueToStringAndType(v interface{}) (string, string) {
	// Null values result in a panic so cast them to false
	if v == nil {
		return "false", "bool"
	}

	varType := reflect.TypeOf(v).Kind().String()

	if varType == "map" {
		v, _ := json.Marshal(v)
		return string(v), "json"
	}

	return fmt.Sprint(v), varType
}
