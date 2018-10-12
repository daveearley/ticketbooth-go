package database

import (
	"database/sql"
	"fmt"
	"github.com/daveearley/product/app/models"
	"github.com/daveearley/product/app/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func InitDb() *sql.DB {
	dbStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dbStr)
	utils.CheckErr(err)

	models.RegisterHooks()

	return db
}
