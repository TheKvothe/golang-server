package service

import (
	"errors"
	"math/rand"

	"../entity"
	"../repository"
)

type service struct{}
type Components struct {
	Components []Component `json:"types"`
}
type Component struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

var (
	repo           repository.PostRepository
	componentsRepo repository.ComponentsRepository = repository.NewComponentsRepository()
)

func NewPostServiceImp(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*service) FindAllType() ([]entity.ComponentType, error) {

	return componentsRepo.FindAllTypes()
}
