package database

import (
	"fmt"
	"github.com/daveearley/product/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres"
)

func InitDb() *gorm.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("postgres", dbinfo)
	utils.CheckErr(err)

	return db
}
