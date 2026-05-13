package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex;not null"`
	Resource string `gorm:"not null"`
	Action   string `gorm:"not null"`
}
