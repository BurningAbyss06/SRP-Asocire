package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Pays     []Pay
	Salt     int `gorm:"not null"`
}

type NewUser struct {
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
}
