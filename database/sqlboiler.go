package database

import (
	"database/sql"
	"fmt"
	"github.com/daveearley/product/app/models/hooks"
	"github.com/daveearley/product/app/utils"
	"github.com/daveearley/product/configs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDb(config *configs.Config) *sql.DB {
	dbStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		config.DbUser,
		config.DbPassword,
		config.DbName,
	)

	db, err := sql.Open(config.DbDriver, dbStr)
	utils.CheckErr(err)

	hooks.Register()

	return db
}
