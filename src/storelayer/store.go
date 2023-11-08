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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS animals (id int NOT NULL AUTO_INCREMENT PRIMARY KEY,name varchar(50),age int)")
	// TODO implement migrate
	return &store{db: db}
}
