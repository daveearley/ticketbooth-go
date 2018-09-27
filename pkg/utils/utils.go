package utils

import (
	"fmt"
	"strconv"
)

// Covert a string to an unsigned int
func Str2Uint(str string) {
	i, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		fmt.Printf("Type: %T \n", i)
		fmt.Println(i)
	}
}
