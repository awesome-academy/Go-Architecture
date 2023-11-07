package httplayer

import (
	"Go-Architecture/src/applayer"
	"github.com/gorilla/mux"
	"net/http"
)

var r *mux.Router

type httpApi struct {
	app applayer.App
}

func New(applayer applayer.App) *httpApi {
	app := &httpApi{
		app: applayer,
	}
	app.setupRoutes()
	return app
}

func (api httpApi) setupRoutes() {
	r = mux.NewRouter()
	r.HandleFunc("/animal", api.createAnimal).Methods("POST")
	r.HandleFunc("/animals", api.getAnimalById).Queries("id", "{id}").Methods("GET")
	r.HandleFunc("/animals", api.getAllAnimals).Methods("GET")
}

func (api *httpApi) Engage() {
	http.ListenAndServe(":3000", r)
}
