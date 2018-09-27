package model

type User struct {
	Model
	Email     string
	Password  string
	FirstName string
	LastName  string
	Status    string
	AccountId int
}

//func (u *User) BeforeSave() (err error)  {
//	bytePassword := []byte(u.Password)
//	u.Password, _ = bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//}
