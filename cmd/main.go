package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"

	"small-crud/internal/controller"
	"small-crud/internal/repositories/sqlite"
)

func main() {
	r := gin.Default()

	database, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	defer database.Close()

	articleRepository := sqlite.NewSqliteRepository(database)

	httpController := controller.NewHttpController(r, articleRepository)
	httpController.Init()

	r.Run()
}
