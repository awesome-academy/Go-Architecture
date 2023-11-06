package animal

import (
	"Go-Architecture/entities"
	"database/sql"
)

type AnimalStorer struct {
	db *sql.DB
}

func New(db *sql.DB) AnimalStorer {
	return AnimalStorer{db: db}
}

func (a AnimalStorer) Create(animal entities.Animal) (entities.Animal, error) {
	res, err := a.db.Exec("INSERT INTO animals (name,age) VALUES(?,?)", animal.Name, animal.Age)

	if err != nil {
		return entities.Animal{}, err
	}

	id, _ := res.LastInsertId()
	animal.ID = int(id)

	return animal, nil
}

func (a AnimalStorer) Get(id int) ([]entities.Animal, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != 0 {
		rows, err = a.db.Query("SELECT * FROM animals where id = ?", id)
	} else {
		rows, err = a.db.Query("SELECT * FROM animals")
	}

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
		_ = rows.Scan(&a.ID, &a.Name, &a.Age)
		animals = append(animals, a)
	}
	return animals, nil
}
