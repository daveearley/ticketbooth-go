package main

import (
	"fmt"
	"github.com/daveearley/product/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres"
)

func main() {
	server := gin.Default()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("postgres", dbinfo)
	checkErr(err)

	defer db.Close()

	api.RegisterRoutes(server, db)

	server.Run(":1234")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
