package repositories

import (
	"database/sql"
	"errors"
	"ramori/internal/models"
)

type UserRepository interface {
	Create(user models.User, password string) (int, error)
	Update(user models.User) (int, error)
	UpdatePassword(user models.User, password string) (int, error)
	Delete(id int) (int, error)
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
			if err := tx.Rollback(); err != nil {
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
			}
		} else {
			if err := tx.Commit(); err != nil {
			}
		}
		if err != nil {
			return
		}
	}()

	if err != nil {
		return 0, err
	}

	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	result, err := tx.Exec(query, user.FirstName, user.LastName, user.Email, password)
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) Update(user models.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
			}
		} else {
			if err := tx.Commit(); err != nil {
			}
		}
	}()
	if err != nil {
		return 0, err
	}

	query := "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	_, err = tx.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return 0, err
	}

	id := user.ID

	return id, nil
}

func (r *MySQLUserRepository) Delete(id int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
			}
		} else {
			if err := tx.Commit(); err != nil {
			}
		}
	}()

	if err != nil {
		return 0, err
	}

	query := "DELETE users FROM users WHERE id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *MySQLUserRepository) UpdatePassword(user models.User, password string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
			}
		} else {
			if err := tx.Commit(); err != nil {
			}
		}
	}()

	if err != nil {
		return 0, err
	}

	query := "UPDATE users SET password = ? WHERE id = ?"
	_, err = tx.Exec(query, password, user.ID)
	if err != nil {
		return 0, err
	}

	id := user.ID

	return id, nil

}
