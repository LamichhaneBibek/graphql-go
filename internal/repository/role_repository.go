package repository

import (
	"errors"

	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(id uint) (*models.Role, error)
	FindByName(name string) (*models.Role, error)
	FindAll() ([]models.Role, error)
	Create(role *models.Role) error
	HasPermission(roleID uint, resource, action string) (bool, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.Preload("Permissions").First(&role, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("role not found")
	}
	return &role, err
}

func (r *roleRepository) FindByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Preload("Permissions").Where("name = ?", name).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("role not found")
	}
	return &role, err
}

func (r *roleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) HasPermission(roleID uint, resource, action string) (bool, error) {
	var count int64
	err := r.db.Table("role_permissions").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ? AND permissions.resource = ? AND permissions.action = ? AND permissions.deleted_at IS NULL", roleID, resource, action).
		Count(&count).Error
	return count > 0, err
}
