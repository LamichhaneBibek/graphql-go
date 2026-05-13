package service

import (
	"errors"

	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"github.com/LamichhaneBibek/graphql-go/internal/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	UpdateRole(userID, roleID uint) (*models.User, error)
	HasRole(userID uint, roleName string) (bool, error)
}

type userService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{userRepo: userRepo, roleRepo: roleRepo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateRole(userID, roleID uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if _, err := s.roleRepo.FindByID(roleID); err != nil {
		return nil, errors.New("role not found")
	}

	user.RoleID = &roleID
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return s.userRepo.FindByID(userID)
}

func (s *userService) HasRole(userID uint, roleName string) (bool, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return false, err
	}
	if user.Role == nil {
		return false, nil
	}
	return user.Role.Name == roleName, nil
}
