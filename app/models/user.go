package models

import "github.com/volatiletech/null"

type User struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Email     string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	Password  string    `boil:"password" json:"password" toml:"password" yaml:"password"`
	FirstName string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName  string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Status    string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	AccountID int       `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Account   *Account  `json:"account"`
	Events    *[]Event  `json:"events"`
}
