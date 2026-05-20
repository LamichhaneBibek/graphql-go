package service

import (
	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"github.com/LamichhaneBibek/graphql-go/internal/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
