package model

import (
	"github.com/daveearley/product/pkg/utils"
)

type User struct {
	Model
	Email     string `json:"email" binding:"email,required"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
	AccountId uint64 `json:"account_id"`
}

func (u *User) BeforeSave() (err error) {
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	return
}
