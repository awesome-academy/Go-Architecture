package main

import (
	"Go-Architecture/datastore/animal"
	handlerAnimal "Go-Architecture/delivery/animal"
	"Go-Architecture/driver"
	"Go-Architecture/util"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	util.LoadEnv()
	// get the mysql configs from env:
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}
	var err error

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}

	datastore := animal.New(db)
	handler := handlerAnimal.New(datastore)

	http.HandleFunc("/animal", handler.Handler)
	fmt.Println(http.ListenAndServe(":9000", nil))
}
