package models

import "gorm.io/gorm"

type Apt_Casa struct {
	gorm.Model
	AddressID uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	Address   Address
	User      User
}
