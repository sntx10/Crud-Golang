package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"small-crud/internal/models"
	"small-crud/internal/repositories"
)

type ArticleController struct {
	r    *gin.RouterGroup
	repo repositories.ArticleRepositoryInterface
}

func NewArticleController(r *gin.RouterGroup, repo repositories.ArticleRepositoryInterface) *ArticleController {
	return &ArticleController{r: r, repo: repo}
}

func (a *ArticleController) Init() {
	a.r.GET("/", a.FindArticles)
	a.r.GET("/:id", a.FindArticle)
	a.r.POST("/", a.CreateArticle)
	a.r.PUT("/:id", a.UpdateArticle)
	a.r.DELETE("/:id", a.DeleteArticle)
}

func (a *ArticleController) FindArticles(c *gin.Context) {
	articles, err := a.repo.FindAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func (a *ArticleController) FindArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	article, err := a.repo.FindArticleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func (a *ArticleController) CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{Title: input.Title, Content: input.Content}
	id, err := a.repo.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	article.Id = int(id)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func (a *ArticleController) UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{Id: id, Title: input.Title, Content: input.Content}
	if err := a.repo.UpdateArticle(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func (a *ArticleController) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := a.repo.DeleteArticle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
