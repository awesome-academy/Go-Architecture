package applayer

import (
	"Go-Architecture/src/httplayer/entities"
	"Go-Architecture/src/storelayer"
)

type App interface {
	CreateAnimal(entities.Animal) (entities.Animal, error)

	GetAnimalById(id int) (entities.Animal, error)

	GetAllAnimals() ([]entities.Animal, error)
}

type app struct {
	store storelayer.Store
}

func New(store storelayer.Store) *app {
	return &app{
		store: store,
	}
}
