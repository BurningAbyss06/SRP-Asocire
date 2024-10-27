package models

type Address struct {
	ID         uint   `gorm:"not null"`
	Is_Apt     bool   `gorm:"default:false"`
	Is_House   bool   `gorm:"default:false"`
	Location   string `gorm:"not null"`
	Name_House string `gorm:"validate:required_if=Is_House true"`
	No_Apt     string `gorm:"validate:required_if=Is_Apt true"`
}
