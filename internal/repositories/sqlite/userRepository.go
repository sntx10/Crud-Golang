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


func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	stmt, err := r.DB.Prepare("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Id)
	return err
}

func (r *UserRepository) DeleteUser(email, password string) error {
	stmt, err := r.DB.Prepare("DELETE FROM users WHERE email = ? AND password = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, password)
	return err
}
