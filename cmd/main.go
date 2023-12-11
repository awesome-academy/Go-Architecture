package main

import (
	"Go-Architecture/api/route"
	"Go-Architecture/bootstrap"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	app := bootstrap.App()

	defer app.CloseDBConnection()

	env := app.Env

	timeout := time.Duration(env.ContextTimeOut) * time.Second

	http := gin.Default()

	route.Setup(env, timeout, app.Db, http)

	err := http.Run()
	if err != nil {
		return
	}
}
