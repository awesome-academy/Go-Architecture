package animal

import (
	"Go-Architecture/driver"
	"Go-Architecture/entities"
	"Go-Architecture/util"
	"database/sql"
	"os"
	"reflect"
	"testing"
)

func initializeMySQL(t *testing.T) *sql.DB {
	util.LoadEnv()
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("Could no connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T) {
	db := initializeMySQL(t)
	animalStore := New(db)
	testAnimalStorer_Create(t, animalStore)
	testAnimalStorer_Get(t, animalStore)
}

func testAnimalStorer_Create(t *testing.T, db AnimalStorer) {
	testcases := []struct {
		req      entities.Animal
		response entities.Animal
	}{
		{req: entities.Animal{Name: "Hen", Age: 1}, response: entities.Animal{ID: 3, Name: "Hen", Age: 1}},
		{req: entities.Animal{Name: "Pig", Age: 2}, response: entities.Animal{ID: 4, Name: "Pig", Age: 2}},
	}

	for i, testcase := range testcases {
		resp, _ := db.Create(testcase.req)

		if !reflect.DeepEqual(resp, testcase.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, testcase.response)
		}
	}
}

func testAnimalStorer_Get(t *testing.T, db AnimalStorer) {
	testcases := []struct {
		id       int
		response []entities.Animal
	}{
		{0, []entities.Animal{{1, "Hippo", 10}, {2, "Ele", 20}}},
		{1, []entities.Animal{{1, "Hippo", 10}}},
	}

	for i, testcase := range testcases {
		resp, _ := db.Get(testcase.id)

		if !reflect.DeepEqual(resp, testcase.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, testcase.response)
		}
	}
}
