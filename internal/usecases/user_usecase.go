package usecases

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ramori/internal/models"
	"ramori/internal/repositories"
)

type UserUseCase interface {
	CreateUser(firstName, lastName, email, password string) (id int, err error)
	Update(id int, firstName, lastName, email string) (int, error)
	UpdatePassword(id int, password string) (int, error)
	Delete(id int) (int, error)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := uc.UserRepo.Create(user, string(hashedPassword))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *userUseCase) Update(id int, firstName, lastName, email string) (int, error) {
	if id == 0 || firstName == "" || lastName == "" || email == "" {
		return 0, errors.New("first name, last name, id and email cannot be empty")
	}

	user := models.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	id, err := uc.UserRepo.Update(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *userUseCase) UpdatePassword(id int, password string) (int, error) {
	if id == 0 || password == "" {
		return 0, errors.New("id and password cannot be empty")
	}

	user := models.User{
		ID: id,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	id, err = uc.UserRepo.UpdatePassword(user, string(hashedPassword))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *userUseCase) Delete(id int) (int, error) {
	if id == 0 {
		return 0, errors.New("id cannot be empty")
	}

	id, err := uc.UserRepo.Delete(id)

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
