package model

type Account struct {
	Model
	Email string `json:"email" binding:"required,email" form:"email"`
	Users []User `json:"users"`
}
