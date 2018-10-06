package database

import (
	"fmt"
	"github.com/daveearley/product/pkg/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"database/sql"
)

const (
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres"
)

func InitDb() *sql.DB {
	dbStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbStr)
	utils.CheckErr(err)

	return db
}