package service

import (
	"errors"

	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"github.com/LamichhaneBibek/graphql-go/internal/repository"
)

type PostService interface {
	GetAll() ([]models.Post, error)
	GetByID(id uint) (*models.Post, error)
	Create(authorID uint, title, content string) (*models.Post, error)
	Update(id, userID uint, title, content *string) (*models.Post, error)
	Delete(id, userID uint) error
	Publish(id uint) (*models.Post, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) GetAll() ([]models.Post, error) {
	return s.postRepo.FindAll()
}

func (s *postService) GetByID(id uint) (*models.Post, error) {
	return s.postRepo.FindByID(id)
}

func (s *postService) Create(authorID uint, title, content string) (*models.Post, error) {
	post := &models.Post{Title: title, Content: content, AuthorID: authorID}
	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}
	return s.postRepo.FindByID(post.ID)
}

func (s *postService) Update(id, userID uint, title, content *string) (*models.Post, error) {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if post.AuthorID != userID {
		return nil, errors.New("forbidden: not the author")
	}

	if title != nil {
		post.Title = *title
	}
	if content != nil {
		post.Content = *content
	}

	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) Delete(id, userID uint) error {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		return errors.New("post not found")
	}

	if post.AuthorID != userID {
		return errors.New("forbidden: not the author")
	}

	return s.postRepo.Delete(id)
}

func (s *postService) Publish(id uint) (*models.Post, error) {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}

	post.Published = true
	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}
