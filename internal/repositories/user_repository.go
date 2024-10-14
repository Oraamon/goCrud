package repositories

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ramori/internal/models"
)

type UserRepository interface {
	Create(user models.User, password string) (int, error)
	GetByID(id int) (*models.User, error)
}

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user models.User, password string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	result, err := tx.Exec(query, user.FirstName, user.LastName, user.Email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *MySQLUserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
