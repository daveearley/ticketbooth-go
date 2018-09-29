package main

import (
	"github.com/daveearley/product/pkg/api"
	"github.com/daveearley/product/pkg/database"
	"github.com/daveearley/product/pkg/model"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := database.InitDb()
	db.AutoMigrate(&model.Account{}, &model.User{})
	defer db.Close()

	api.RegisterRoutes(server, db)

	server.Run(":1234")
}
