package storelayer

import (
	"Go-Architecture/src/httplayer/entities"
	"database/sql"
)

func (s *store) CreateAnimal(animal entities.Animal) (entities.Animal, error) {
	res, err := s.db.Exec("INSERT INTO animals (name,age) VALUES(?,?)", animal.Name, animal.Age)

	if err != nil {
		return entities.Animal{}, err
	}

	id, _ := res.LastInsertId()
	animal.ID = int(id)

	return animal, nil
}

func (s *store) GetAllAnimals() ([]entities.Animal, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = s.db.Query("SELECT * FROM animals")

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// No-handle
		}
	}(rows)

	var animals []entities.Animal

	for rows.Next() {
		var a entities.Animal
		err = rows.Scan(&a.ID, &a.Name, &a.Age)
		if err != nil {
			continue
		}
		animals = append(animals, a)
	}
	return animals, nil
}

func (s *store) GetAnimalById(id int) (entities.Animal, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = s.db.Query("SELECT * FROM animals where id = ?", id)

	if err != nil {
		return entities.Animal{}, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// No-handle
		}
	}(rows)

	var animal entities.Animal

	for rows.Next() {
		err = rows.Scan(&animal.ID, &animal.Name, &animal.Age)
	}
	if err != nil {
		return entities.Animal{}, err
	}

	return animal, nil
}
