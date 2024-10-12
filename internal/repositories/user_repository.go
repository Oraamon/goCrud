package repositories

import (
	"database/sql"
	"errors"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Create(user User) (int, error)
	GetByID(id int) (*User, error)
}

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user User) (int, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *MySQLUserRepository) GetByID(id int) (*User, error) {
	var user User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
