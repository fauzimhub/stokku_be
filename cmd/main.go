package main

import (
	"time"

	"github.com/fauzimhub/stokku/api/route"
	"github.com/fauzimhub/stokku/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(timeout, db, gin)

	gin.Run(env.ServerAddress)
}
