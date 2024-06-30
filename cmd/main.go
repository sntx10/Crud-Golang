package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"

	"small-crud/internal/controllers"
	"small-crud/internal/middlewares"

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
	userRepo := sqlite.NewUserRepository(database)

	authController := controllers.NewAuthController(userRepo)
	authController.Init(r)

	articleRoutes := r.Group("/articles")
	articleRoutes.Use(middlewares.AuthMiddleware())
	{
		httpController := controllers.NewHttpController(articleRoutes, articleRepository)
		httpController.Init()
	}

	r.Run()
}
