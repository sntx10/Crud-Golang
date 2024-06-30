package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"small-crud/internal/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	r  *gin.Engine
	db *sql.DB
}

func NewArticleController(r *gin.Engine, db *sql.DB) *ArticleController {
	return &ArticleController{r: r, db: db}
}

func (a *ArticleController) Init() {
	a.r.GET("/articles", a.FindArticles)
	a.r.GET("/articles/:id", a.FindArticle)
	a.r.POST("/articles", a.CreateArticle)
	a.r.PUT("/articles/:id", a.UpdateArticle)
	a.r.DELETE("/articles/:id", a.DeleteArticle)
}

func (a *ArticleController) FindArticles(c *gin.Context) {
	var articles []models.Article
	var totalRecords int64

	title := c.DefaultQuery("title", "")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	row := a.db.QueryRow("SELECT COUNT(*) FROM articles")
	err := row.Scan(&totalRecords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := "SELECT id, title, content FROM articles"
	if title != "" {
		query += " WHERE title LIKE ?"
		query += " LIMIT ? OFFSET ?"
		rows, err := a.db.Query(query, "%"+title+"%", limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var article models.Article
			if err := rows.Scan(&article.Id, &article.Title, &article.Content); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			articles = append(articles, article)
		}
	} else {
		query += " LIMIT ? OFFSET ?"
		rows, err := a.db.Query(query, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var article models.Article
			if err := rows.Scan(&article.Id, &article.Title, &article.Content); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			articles = append(articles, article)
		}
	}

	pagination := map[string]interface{}{
		"total_records": totalRecords,
		"total_pages":   (totalRecords + int64(limit) - 1) / int64(limit),
		"current_page":  page,
		"limit":         limit,
	}

	c.JSON(http.StatusOK, gin.H{"data": articles, "pagination": pagination})
}

func (a *ArticleController) FindArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var article models.Article
	row := a.db.QueryRow("SELECT id, title, content FROM articles WHERE id = ?", id)
	err = row.Scan(&article.Id, &article.Title, &article.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
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

	result, err := a.db.Exec("INSERT INTO articles (title, content) VALUES (?, ?)", input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{Id: int(id), Title: input.Title, Content: input.Content}
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

	_, err = a.db.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?", input.Title, input.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (a *ArticleController) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	_, err := a.db.Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
