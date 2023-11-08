package main

import "Go-Architecture/bootstrap"

func main() {
	app := bootstrap.App()

	defer app.CloseDBConnection()

}
