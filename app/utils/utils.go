package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"github.com/volatiletech/sqlboiler/types"
	dec "github.com/ericlagergren/decimal"
)

func IntToDecimal(v int64) types.Decimal {
	return types.NewDecimal(dec.New(v,2))
}

// Covert a string to an int
func Str2int(str string) int {
	i, err := strconv.Atoi(str)
	if err == nil {
		return i
	}

	return 1
}

// HashPassword returns a bcrypt hash of a given string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash validates a given password string matches its hashed counterpart
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckErr checks for an error and panics if found
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
