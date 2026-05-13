package repository

import (
	"errors"

	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindByID(id uint) (*models.Post, error)
	FindAll() ([]models.Post, error)
	Create(post *models.Post) error
	Update(post *models.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("Author.Role.Permissions").First(&post, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("post not found")
	}
	return &post, err
}

func (r *postRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("Author.Role.Permissions").Find(&posts).Error
	return posts, err
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
