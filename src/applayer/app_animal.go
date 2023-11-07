package applayer

import (
	"Go-Architecture/src/httplayer/entities"
)

func (app app) CreateAnimal(animal entities.Animal) (entities.Animal, error) {
	return app.store.CreateAnimal(animal)
}

func (app app) GetAnimalById(id int) (entities.Animal, error) {
	return app.store.GetAnimalById(id)
}

func (app app) GetAllAnimals() ([]entities.Animal, error) {
	return app.store.GetAllAnimals()
}
