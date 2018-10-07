package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// Covert a string to an int
func Str2int(str string) int {
	i, err := strconv.Atoi(str)
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
