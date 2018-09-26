package model

type User struct {
	Model
	Email     string
	Password  string
	FirstName string
	LastName  string
	Status    string
}
