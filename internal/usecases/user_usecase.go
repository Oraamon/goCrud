package usecases

import (
	"errors"
	"ramori/internal/models"
	"ramori/internal/repositories"
)

type UserUseCase interface {
	CreateUser(firstName, lastName, email, password string) (id int, err error)
	GetUserByID(id int) (*models.User, error)
}

type userUseCase struct {
	UserRepo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return &userUseCase{
		UserRepo: repo,
	}
}

func (uc *userUseCase) CreateUser(firstName, lastName, email, password string) (int, error) {
	if firstName == "" || lastName == "" || email == "" || password == "" {
		return 0, errors.New("first name, last name, email, and password cannot be empty")
	}

	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	id, err := uc.UserRepo.Create(user, password)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *userUseCase) GetUserByID(id int) (*models.User, error) {
	user, err := uc.UserRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
