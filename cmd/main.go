package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	_ "modernc.org/sqlite"

	"small-crud/internal/controller"
	"small-crud/internal/models"
)

func main() {
	r := gin.Default()
	
	database, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	defer database.Close()

	models.ConnectDatabase(database)

	httpController := controller.NewHttpController(r, database)
	httpController.Init()

	r.Run()
}
