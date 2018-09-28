package utils

import (
	"strconv"
)

// Covert a string to an unsigned int
func Str2Uint(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		return i
	}

	return 1
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
