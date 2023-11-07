package storelayer

import (
	"Go-Architecture/src/httplayer/entities"
	"Go-Architecture/src/storelayer/driver"
	"Go-Architecture/src/utils"
	"database/sql"
	"os"
)

type Store interface {
	CreateAnimal(entities.Animal) (entities.Animal, error)

	GetAnimalById(id int) (entities.Animal, error)

	GetAllAnimals() ([]entities.Animal, error)
}

type store struct {
	db *sql.DB
}

func New() *store {
	utils.LoadEnv()
	// get the mysql configs from env:
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		panic("Failed to connect database!")
	}
	// TODO implement db.migrate
	return &store{db: db}
}
