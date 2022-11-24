package in_memory

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type UserRepository struct {
	user repositories.IUser
}

func (u UserRepository) GetUserByUsername(username string) (models.User, error) {
	return models.User{
		ID:                  1,
		Name:                username,
		Email:               username + "@gmail.com",
		Reports:             0,
		RemainingReports:    3,
		Role:                "Пользователь",
		UnbanDate:           "01.01.1970",
		TransformedPassword: "21jif90iKSDJ9214l",
	}, nil
}

func (u UserRepository) GetUserByEmail(email string) (models.User, error) {
	return models.User{
		ID:                  1,
		Name:                "Alex",
		Email:               email,
		Reports:             0,
		RemainingReports:    3,
		Role:                "Пользователь",
		UnbanDate:           "01.01.1970",
		TransformedPassword: "21jif90iKSDJ9214l",
	}, nil
}

func (u UserRepository) Create(user *models.User) (err error) {
	return err
}

func (u UserRepository) Ban(user *models.User) (err error) {
	return err
}

func (u UserRepository) Delete(user *models.User) (err error) {
	return err
}
