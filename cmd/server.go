package main

import (
	"github.com/daveearley/product/pkg/api"
	"github.com/daveearley/product/pkg/database"
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	server := gin.Default()

	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	db := database.InitDb()
	defer db.Close()

	api.RegisterRoutes(server, db)

	server.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
