package animal

import (
	"Go-Architecture/datastore"
	"Go-Architecture/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type AnimalHandler struct {
	datastore datastore.Animal
}

func New(animal datastore.Animal) AnimalHandler {
	return AnimalHandler{animal}
}

func (animalHandler AnimalHandler) Handler(w *httptest.ResponseRecorder, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		animalHandler.get(w, r)
	case http.MethodPost:
		animalHandler.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (animalHandler AnimalHandler) get(w *httptest.ResponseRecorder, r *http.Request) {
	id := r.URL.Query().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("Invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := animalHandler.datastore.Get(i)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve animal!"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}

func (animalHandler AnimalHandler) create(w *httptest.ResponseRecorder, r *http.Request) {
	var animal entities.Animal

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &animal)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := animalHandler.datastore.Create(animal)
	if err != nil {
		_, _ = w.Write([]byte("could not create animal"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}
