package model

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	UserID      uint       `gorm:"not null"`
	FirstName   string     `gorm:"type:varchar(100)"`
	LastName    string     `gorm:"type:varchar(100)"`
	Email       string     `gorm:"type:varchar(256)"`
	PhoneNumber string     `gorm:"type:varchar(32)"`
	PhotoUrl    string     `gorm:"type:varchar(256)"`
	Description string     `gorm:"type:varchar(512)"`
	City        string     `gorm:"type:varchar(100)"`
	State       string     `gorm:"type:varchar(2)"`
	ZipCode     string     `gorm:"type:varchar(32)"`
	Categories  []Category `gorm:"many2many:contact_categories;"`
}
