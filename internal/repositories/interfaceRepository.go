package repositories

import "small-crud/internal/models"

type ArticleRepositoryInterface interface {
    FindArticleById(int) (models.Article, error)
    FindAllArticles() ([]models.Article, error)
    CreateArticle(models.Article) (int64, error)
    UpdateArticle(models.Article) error
    DeleteArticle(int) error
}

type UserRepositoryInterface interface{
	CreateUser(user models.User) (int64, error)
	FindUserByEmail(email string) (models.User, error)
}