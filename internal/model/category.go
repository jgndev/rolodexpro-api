package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	Name        string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:varchar(512)"`
	Contacts    []Contact `gorm:"many2many:contact_categories;"`
}
