package main

import (
	"Go-Architecture/api/route"
	"Go-Architecture/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	defer app.CloseDBConnection()

	env := app.Env

	timeout := time.Duration(env.ContextTimeOut) * time.Second

	http := gin.Default()

	route.Setup(env, timeout, app.Db, http)
}
