package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, name string, email string) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)", name, email, time.Now())
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
