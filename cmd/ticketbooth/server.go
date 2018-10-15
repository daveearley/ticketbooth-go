package main

import (
	"github.com/daveearley/product/app/api/routes"
	"github.com/daveearley/product/database"
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"os"
)

func main() {
	server := gin.Default()

	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	if os.Getenv("APP_DEBUG") == "true" {
		boil.DebugMode = true
	}

	db := database.InitDb()
	defer db.Close()

	routes.Register(server, db)

	server.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
