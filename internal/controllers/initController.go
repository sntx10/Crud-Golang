package controllers

import (
	"github.com/gin-gonic/gin"
	"small-crud/internal/repositories"
)

type HttpController struct {
	DB repositories.ArticleRepositoryInterface
	R  *gin.RouterGroup
}

func NewHttpController(r *gin.RouterGroup, db repositories.ArticleRepositoryInterface) *HttpController {
	return &HttpController{DB: db, R: r}
}

func (h *HttpController) Init() {
	articleController := NewArticleController(h.R, h.DB)
	articleController.Init()
}
