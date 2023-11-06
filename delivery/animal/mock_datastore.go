package animal

import (
	"Go-Architecture/entities"
	"errors"
)

type mockDatastore struct{}

func (m mockDatastore) Get(id int) ([]entities.Animal, error) {
	if id == 1 {
		return nil, errors.New("db error")
	} else if id == 2 {
		return []entities.Animal{{2, "Dog", 8}}, nil
	}

	return []entities.Animal{{1, "Ken", 23}, {2, "Dog", 8}}, nil
}

func (m mockDatastore) Create(animal entities.Animal) (entities.Animal, error) {
	if animal.Age == 12 {
		return entities.Animal{}, errors.New("db error")
	}

	return entities.Animal{12, "Maggie", 10}, nil
}
