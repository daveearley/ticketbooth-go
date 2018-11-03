package database

import (
	"database/sql"
	"fmt"
	"github.com/daveearley/ticketbooth/app/models/modelhooks"
	"github.com/daveearley/ticketbooth/app/utils"
	"github.com/daveearley/ticketbooth/configs"
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

	modelhooks.Register()

	return db
}
