package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(db *sql.DB, firstName, lastName, email, password string) (int64, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO users (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?)", firstName, lastName, email, hashedPassword, time.Now())
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, first_name, last_name, email, created_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
