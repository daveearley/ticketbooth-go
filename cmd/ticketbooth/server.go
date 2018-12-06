package main

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api"
	"github.com/daveearley/ticketbooth/configs"
	"github.com/daveearley/ticketbooth/database"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	binding.Validator = new(app.DefaultValidator)
	server := gin.Default()
	config := configs.LoadConfig()

	if config.AppDebug {
		boil.DebugMode = true
	}

	db := database.InitDb(config)

	defer db.Close()

	api.BootstrapAndRegisterRoutes(server, db, config)

	server.Run(config.AppHost + ":" + config.AppPort)
}
