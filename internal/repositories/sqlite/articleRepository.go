package sqlite

import (
	"database/sql"
	"small-crud/internal/models"
	"small-crud/internal/repositories"
)

type SqliteRepository struct {
	DB *sql.DB
}


func NewSqliteRepository(db *sql.DB) repositories.ArticleRepositoryInterface {
	return &SqliteRepository{DB: db}
}

func (s *SqliteRepository) FindArticleById(id int) (models.Article, error) {
	var article models.Article
	row := s.DB.QueryRow("SELECT id, title, content, date FROM articles WHERE id = ?", id)
	err := row.Scan(&article.Id, &article.Title, &article.Content)

	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *SqliteRepository) FindAllArticles() ([]models.Article, error) {
	var articles []models.Article
	rows, err := s.DB.Query("SELECT id, title, content FROM articles")
	if err != nil {
		return articles, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Content); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (s *SqliteRepository) CreateArticle(article models.Article) (int64, error) {
	result, err := s.DB.Exec("INSERT INTO articles (title, content) VALUES (?, ?)", article.Title, article.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SqliteRepository) UpdateArticle(article models.Article) error {
	_, err := s.DB.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?", article.Title, article.Content, article.Id)
	return err
}

func (s *SqliteRepository) DeleteArticle(id int) error {
	_, err := s.DB.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}
