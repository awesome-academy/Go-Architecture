package datastore

import "Go-Architecture/entities"

type Animal interface {
	Get(id int) ([]entities.Animal, error)

	Create(entities.Animal) (entities.Animal, error)
}
