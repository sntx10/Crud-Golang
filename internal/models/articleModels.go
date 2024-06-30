package models

import "time"


type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Date time.Time `json:"date"`
	Content string `json:"content"`
}

type CreateArticleInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateArticleInput struct {
	Title string `json:"title"`
	Content string `json:"content"`

}