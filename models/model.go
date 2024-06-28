package models

import "time"


type Article struct {
	Id int `json:"id" gorm:"type:int;primary_key"`
	Title string `json:"title" gorm:"type:varchar(255)"`
	Date time.Time `json:"date" gorm:"type:datetime"`
	Content string `json:"content" gorm:"type:varchar(255)"`
}

type CreateArticleInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateArticleInput struct {
	Title string `json:"title"`
	Content string `json:"content"`

}