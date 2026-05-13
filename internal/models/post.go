package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Published bool   `gorm:"default:false"`
	AuthorID  uint   `gorm:"not null"`
	Author    User
}
