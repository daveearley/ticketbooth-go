package main

import (
	"github.com/daveearley/product/app/api/routes"
	"github.com/daveearley/product/configs"
	"github.com/daveearley/product/database"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	server := gin.Default()
	config := configs.LoadConfig()

	if config.AppDebug {
		boil.DebugMode = true
	}

	db := database.InitDb(config)
	defer db.Close()

	routes.Register(server, db, config)

	server.Run(config.AppHost + ":" + config.AppPort)
}
