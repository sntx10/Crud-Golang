package sqlite

import (
	"database/sql"
	"small-crud/internal/models"
	"small-crud/internal/repositories"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	stmt, err := r.DB.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *UserRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	row := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	return user, err
}
