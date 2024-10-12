package usecases

import (
	"errors"
	"ramori/internal/models"
	"ramori/internal/repositories"
)

type UserUseCase interface {
	CreateUser(name, email string) (id int, err error)
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

func (uc *userUseCase) CreateUser(name, email string) (id int, err error) {

	if name == "" || email == "" {
		return 0, errors.New("name and email cannot be empty")
	}

	user := models.User{
		Name:  name,
		Email: email,
	}

	id, err = uc.UserRepo.Create(user)
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
