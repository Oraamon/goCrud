package usecases

import (
	"errors"
	"ramori/internal/models"
	"ramori/internal/repositories"
)

type UserUseCase struct {
	UserRepo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
	}
}

func (uc *UserUseCase) CreateUser(name, email string) (models.User, error) {

	if name == "" || email == "" {
		return models.User{}, errors.New("name and email cannot be empty")
	}

	user := models.User{
		Name:  name,
		Email: email,
	}
	id, err := uc.UserRepo.Create(user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = id
	return user, nil
}

func (uc *UserUseCase) GetUserByID(id int) (models.User, error) {
	user, err := uc.UserRepo.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}
