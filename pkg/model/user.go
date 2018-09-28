package model

type User struct {
	Model
	Email     string `json:"email" binding:"email,required"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
	AccountId int    `json:"account_id"`
}

//func (u *User) BeforeSave() (err error)  {
//	bytePassword := []byte(u.Password)
//	u.Password, _ = bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//}
