package hooks

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/utils"
	"github.com/volatiletech/sqlboiler/boil"
)

// Before insert hook to hash user password
func userInsertHook(executor boil.Executor, u *models.User) error {
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword

	return nil
}

func Register() {
	models.AddUserHook(boil.BeforeInsertHook, userInsertHook)
}
