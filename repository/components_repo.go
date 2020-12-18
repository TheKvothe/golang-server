package repository

import "../entity"

type ComponentsRepository interface {
	FindAll() ([]entity.ComponentType, error)
}
