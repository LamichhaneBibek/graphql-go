package seed

import (
	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"gorm.io/gorm"
)

func Roles(db *gorm.DB) {
	for _, name := range []string{"admin", "user"} {
		db.FirstOrCreate(&models.Role{}, models.Role{Name: name})
	}
}
