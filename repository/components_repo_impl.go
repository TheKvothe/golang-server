package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"../entity"
)

type repos struct{}
type Components struct {
	Components []Component `json:"types"`
}
type Component struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

//NewComponentsRepository creates a new repo
func NewComponentsRepository() ComponentsRepository {
	return &repos{}
}

func (*repos) FindAllTypes() ([]entity.ComponentType, error) {
	var types []entity.ComponentType
	file, err := os.Open("/Users/arthurbernal/dev/golang-clean/config/ComponentType.json")
	_ = err
	//fmt.Print(file)
	input, err := ioutil.ReadAll(file)
	var components Components
	json.Unmarshal([]byte(input), &components)
	//fmt.Println(components)
	for i := 0; i < len(components.Components); i++ {
		componentType := entity.ComponentType{
			ID:          components.Components[i].ID,
			Name:        components.Components[i].Name,
			Description: components.Components[i].Description,
		}
		types = append(types, componentType)
	}
	return types, nil
}
