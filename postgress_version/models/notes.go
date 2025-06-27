package models

import "gorm.io/gorm"

type Notes struct {
	ID      uint   `gorm:"primary key;autoIncrement" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Creator string `json:"creator"`
}

func MigrateNotes(db *gorm.DB) error {
	err := db.AutoMigrate(&Notes{})
	return err
}
