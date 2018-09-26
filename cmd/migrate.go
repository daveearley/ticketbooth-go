package main

import (
	"github.com/daveearley/product/pkg/db"
	"github.com/daveearley/product/pkg/model"
)

func main() {
	db := db.InitDb()
	defer db.Close()

	db.AutoMigrate(&model.Account{}, &model.User{})
}
