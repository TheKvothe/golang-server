package factory

import (
	"fmt"

	"../repository"
	"../service"
)

//GetService creates the service depending on the type
func GetService(serviceType string, repository repository.PostRepository) (service.PostService, error) {
	switch serviceType {
	case "implementation":
		return service.NewPostServiceImp(repository), nil
	default:
		return nil, fmt.Errorf("Wrong service type passed")
	}
}
