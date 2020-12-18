package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"../entity"
	"../errors"
	"../factory"
	"../repository"
	"../service"
)

type controller struct{}
type Components struct {
	Components []Component `json:"types"`
}
type Component struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

var (
	postService    service.PostService
	postRepository repository.PostRepository
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	GetTypes(response http.ResponseWriter, request *http.Request)
	GetTypes2(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(repository repository.PostRepository) PostController {
	postRepository = repository
	return &controller{}
}

func setService(serviceType string) {
	aux, err := factory.GetService(serviceType, postRepository)
	//falta tratar bien el error
	if err != nil {
		fmt.Errorf("Wrong service type passed")
	}
	postService = aux
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	//El set service el parametro lo tiene que coger del request
	setService("implementation")
	fmt.Print(request)
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetTypes(response http.ResponseWriter, request *http.Request) {
	//El set service el parametro lo tiene que coger del request
	setService("implementation")
	//fmt.Print(request)
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAllType()
	//fmt.Println(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetTypes2(response http.ResponseWriter, request *http.Request) {
	//El set service el parametro lo tiene que coger del request
	setService("implementation")
	response.Header().Set("Content-Type", "application/json")
	file, err := os.Open("/Users/arthurbernal/dev/golang-clean/config/ComponentType.json")
	_ = err
	//fmt.Print(file)
	input, err := ioutil.ReadAll(file)
	var components Components
	json.Unmarshal([]byte(input), &components)
	fmt.Println(components)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(components)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	//El set service el parametro lo tiene que coger del request
	setService("implementation")
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	fmt.Print(request)
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
