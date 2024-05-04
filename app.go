package main

import (
	"myproject/database"
	// "net/http"
	"myproject/router"
	// "github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDb()
	database.MigrateDb()
}

func main() {
	r := router.SetupRouter()
	r.Run("localhost:9090")
}