package model

import (
	"github.com/jgndev/rolodexpro-api/internal/types"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string         `gorm:"type:varchar(100);unique_index"`
	Password    string         `gorm:"size:255"`
	Salt        string         `gorm:"size:255"`
	DisplayName string         `gorm:"type:varchar(100)"`
	PhotoUrl    string         `gorm:"type:varchar(256)"`
	UserRole    types.UserRole `gorm:"type:int"`
	Auth0ID     string         `gorm:"type:varchar(100);unique_index"`
	Contacts    []Contact      `gorm:"FOREIGNKEY:UserID"`
	Categories  []Category     `gorm:"FOREIGNKEY:UserID"`
}
