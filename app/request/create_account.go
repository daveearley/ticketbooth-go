package request

type CreateAccount struct {
	Email           string `json:"email" binding:"required,email"`
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Password        string `json:"password" binding:"required,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}
