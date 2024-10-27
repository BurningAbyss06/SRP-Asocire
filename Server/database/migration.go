package database

import (
	"log"

	"github.com/BurningAbyss06/SRP-Asocire/Server/models"
	"gorm.io/gorm"
)

func Create_migration(DB *gorm.DB) {

	if err := DB.AutoMigrate(
		models.Address{},
		models.Apt_Casa{},
		models.Pay{},
		models.User{},
	); err != nil {
		log.Fatalln(err)
	}

}
