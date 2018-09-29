package utils

import (
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
