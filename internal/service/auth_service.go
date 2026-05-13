package service

import (
	"errors"

	"github.com/LamichhaneBibek/graphql-go/internal/auth"
	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"github.com/LamichhaneBibek/graphql-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthResult struct {
	Token string
	User  *models.User
}

type AuthService interface {
	Register(name, email, password string) (*AuthResult, error)
	Login(email, password string) (*AuthResult, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(name, email, password string) (*AuthResult, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{Name: name, Email: email, Password: string(hashed)}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResult{Token: token, User: user}, nil
}

func (s *authService) Login(email, password string) (*AuthResult, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResult{Token: token, User: user}, nil
}
