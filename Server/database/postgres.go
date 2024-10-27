package database

import (
	"log"

	"github.com/BurningAbyss06/SRP-Asocire/Server/utilities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	var DB *gorm.DB
	var err error
	DSN := utilities.Lector("CONNSTR")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return DB
}
