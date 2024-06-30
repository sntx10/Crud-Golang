package controller

import (
	"github.com/gin-gonic/gin"
	"database/sql"
)

type HttpController struct {
	DB *sql.DB
	R  *gin.Engine
}

func NewHttpController(r *gin.Engine, db *sql.DB) *HttpController {
	return &HttpController{DB: db, R: r}
}

func (h *HttpController) Init() {
    articleController := NewArticleController(h.R, h.DB)
    articleController.Init()
}
