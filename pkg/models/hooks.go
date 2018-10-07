package models

import (
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/daveearley/product/pkg/utils"
)

// Before insert hook to hash user password
func userInsertHook(executor boil.Executor, u *models.User) error {
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword

	return nil
}

func RegisterHooks()  {
	models.AddUserHook(boil.BeforeInsertHook, userInsertHook)
}
