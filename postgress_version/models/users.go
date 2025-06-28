package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        //shortcut to add ID, timestamps, and soft delete fields to your struct automatically.
	Username   string `gorm:"unique" json="username"`
	Password   string `json"password"`
}
