package controller

import (
	"net/http"
	"strconv"

	"small-crud/models"

	"github.com/gin-gonic/gin"
)


// func FindArticles(c *gin.Context) {
// 	var articles []models.Article
// 	models.DB.Find(&articles)
	
// 	c.JSON(http.StatusOK, gin.H{"data": articles})
// 	}

func FindArticles(c *gin.Context) {
	var articles []models.Article
	var totalRecords int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page -1) * limit
	
	models.DB.Model(&models.Article{}).Count(&totalRecords)

	if err := models.DB.Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagination := map[string]interface{}{
		"total_records": totalRecords,
		"total_pages": (totalRecords + int64(limit) - 1) / int64(limit),
		"current_page": page,
		"limit":		limit,
	}
	
	c.JSON(http.StatusOK, gin.H{"data": articles, "pagination": pagination})
}


func CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create article
	article := models.Article{Title: input.Title, Content: input.Content}
	models.DB.Create(&article)

	c.JSON(http.StatusOK, gin.H{"data": article})
}


func FindArticle(c *gin.Context) {
	var article models.Article
	title := c.Param("title")

	if err := models.DB.Where("title = ?", title).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})

}

func UpdateArticle(c *gin.Context) {
	var article models.Article
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&article).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": article})
}

func DeleteArticle(c *gin.Context) {
	var article models.Article
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&article)

	c.JSON(http.StatusOK, gin.H{"data": true})
}