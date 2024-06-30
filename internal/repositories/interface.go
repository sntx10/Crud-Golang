package repositories

import "small-crud/internal/models"

type ArticleRepositoryInterface interface {
    FindArticleById(int) (models.Article, error)
    FindAllArticles() ([]models.Article, error)
    CreateArticle(models.Article) (int64, error)
    UpdateArticle(models.Article) error
    DeleteArticle(int) error
}
