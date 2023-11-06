package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
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
	_, _ = w.Write([]byte("Dog"))
}
