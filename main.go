package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type mysqlConfig struct {
	host     string
	user     string
	password string
	port     string
	db       string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := mysqlConfig{
		host:     os.Getenv("SQL_HOST"),
		user:     os.Getenv("SQL_USER"),
		password: os.Getenv("SQL_PASSWORD"),
		port:     os.Getenv("SQL_PORT"),
		db:       os.Getenv("SQL_DB"),
	}

	var errConnectDB error
	db, errConnectDB = connectToMySQL(conf)
	if errConnectDB != nil {
		log.Println("could not connect to mysql, err:", err)
		return
	}

	http.HandleFunc("/animal", handler)
	fmt.Println(http.ListenAndServe(":9000", nil))
}

func connectToMySQL(conf mysqlConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.user, conf.password, conf.host, conf.port, conf.db)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		post(w, r)
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		get(w, r)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type Animal struct {
	Name string
	Age  int
}

func post(w http.ResponseWriter, r *http.Request) {
	var animal Animal
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &animal)
	_, err := db.Exec("INSERT INTO animals (Name,Age) VALUES(?,?)", animal.Name, animal.Age)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}

	_, _ = w.Write([]byte("success"))
}

func get(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM animals")

	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Something unexpected happened"))
	}

	defer rows.Close()
	var animals []Animal
	for rows.Next() {
		var animal Animal
		_ = rows.Scan(&animal.Name, &animal.Age)
		animals = append(animals, animal)
	}

	response, _ := json.Marshal(animals)
	_, _ = w.Write(response)
}
