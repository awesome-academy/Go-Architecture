package main

import (
	"bytes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func Test_Handler(t *testing.T) {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}
	// initializing db
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
		t.Errorf("Could not connect to sql, err:%v", errConnectDB)
	}

	testcases := []struct {
		// input
		method string
		body   []byte

		// output
		expectedStatusCode int
		expectedResponse   []byte
	}{
		{"POST", []byte(`[{"Name":"Hippo","Age":10}]`), http.StatusOK, []byte(`success`)},
		{"GET", nil, http.StatusOK, []byte(`[{"Name":"Hippo","Age":10}]`)},
		{"DELETE", nil, http.StatusMethodNotAllowed, nil},
	}

	for _, testcase := range testcases {
		req := httptest.NewRequest(testcase.method, "/animal", bytes.NewReader(testcase.body))
		w := httptest.NewRecorder()

		h := http.HandlerFunc(handler)
		h.ServeHTTP(w, req)

		if w.Code != testcase.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", testcase.expectedStatusCode, w.Code)
		}

		expectedResponse := bytes.NewBuffer(testcase.expectedResponse)
		if !reflect.DeepEqual(w.Body, expectedResponse) {
			t.Errorf("Expected %v\tGot %v", expectedResponse.String(), w.Body.String())
		}
	}
}
