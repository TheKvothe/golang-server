package service

import (
	"../entity"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindAllType() ([]entity.ComponentType, error)
}
