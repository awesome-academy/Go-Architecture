package httplayer

import (
	"Go-Architecture/src/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (api httpApi) CreateAnimal(writer http.ResponseWriter, request *http.Request) {
	var animal entities.Animal

	body, _ := io.ReadAll(request.Body)

	err := json.Unmarshal(body, &animal)
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("invalid body"))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := api.app.CreateAnimal(animal)
	if err != nil {
		_, _ = writer.Write([]byte("Fail to create animal"))
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, err = json.Marshal(resp)
	_, _ = writer.Write(body)
}

func (api httpApi) GetAnimalById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = writer.Write([]byte("Invalid parameter id"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := api.app.GetAnimalById(i)
	if err != nil {
		_, _ = writer.Write([]byte("Could not retrieve animal!"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(resp)
	_, _ = writer.Write(body)
}

func (api httpApi) GetAllAnimals(writer http.ResponseWriter, request *http.Request) {
	resp, err := api.app.GetAllAnimals()
	if err != nil {
		_, _ = writer.Write([]byte("Could not retrieve animals!"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(resp)
	_, _ = writer.Write(body)
}
