package repository

import "../entity"

type ComponentsRepository interface {
	FindAllTypes() ([]entity.ComponentType, error)
}
