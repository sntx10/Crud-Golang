package controller

import (
	"github.com/gin-gonic/gin"
	"small-crud/internal/repositories"
)

type HttpController struct {
	DB repositories.ArticleRepositoryInterface
	R  *gin.Engine
}

func NewHttpController(r *gin.Engine, db repositories.ArticleRepositoryInterface) *HttpController {
	return &HttpController{DB: db, R: r}
}

func (h *HttpController) Init() {
	articleController := NewArticleController(h.R, h.DB)
	articleController.Init()
}
