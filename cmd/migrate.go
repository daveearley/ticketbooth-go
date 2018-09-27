package main

import (
	"github.com/daveearley/product/pkg/database"
	"github.com/daveearley/product/pkg/model"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	db.AutoMigrate(&model.Account{}, &model.User{})
}
