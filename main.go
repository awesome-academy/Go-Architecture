package main

import (
	"Go-Architecture/src/applayer"
	"Go-Architecture/src/httplayer"
	"Go-Architecture/src/storelayer"
)

func main() {
	// create store layer
	storeLayer := storelayer.New()

	// create app layer
	appLayer := applayer.New(storeLayer)

	// create http layer
	api := httplayer.New(appLayer)

	api.Engage()
}
