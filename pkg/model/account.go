package model

type Account struct {
	Model
	Id    uint64 `json:"id" gorm:"primary_key"`
	Email string `json:"email" binding:"required" form:"email"`
	Users []User `json:"users"`
}
