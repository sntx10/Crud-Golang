package main

import (
	"github.com/gin-gonic/gin"

	"small-crud/controller"
	"small-crud/models" //new
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/articles", controller.FindArticles)
	r.GET("/article/:title", controller.FindArticle)
	r.POST("/article", controller.CreateArticle)
	r.PATCH("/article/:id", controller.UpdateArticle)
	r.DELETE("/article/:id", controller.DeleteArticle)
	r.Run()

	err := r.Run()
	if err != nil {
		return
	}
}
