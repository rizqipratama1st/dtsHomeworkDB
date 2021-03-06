package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null;"`
	Slug        string `gorm:"type:varchar(255);unique;not null;"`
	Description string `gorm:"type:text;not null"`
	Duration    uint   `gorm:"type:int(5);not null"`
	Image       string `gorm:"type:varchar(255);not null"`
}

func MigrateMovie(db *gorm.DB) {
	db.AutoMigrate(&Movie{})
}