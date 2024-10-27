package models

import (
	"time"

	"gorm.io/gorm"
)

type Pay struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	No_Bank string `gorm:"not null"`
	Concept string
	Date    time.Time  `gorm:"not null"`
	Month   time.Month `gorm:"not null"`
	Tax     float32    `gorm:"not null"`
	Amount  float32    `gorm:"not null"`
}
