package model

type Account struct {
	Model
	Email string `json:"email" binding:"required" form:"email"`
	Users []User `json:"users"`
}
