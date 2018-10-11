package models

// AuthUser is a populated from JWT claims
type AuthUser struct {
	UserId    int `json:"user_id"`
	AccountId int `json:"account_id"`
}
