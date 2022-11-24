package psql

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type UserRepository struct {
	user repositories.IUser
}

func (u UserRepository) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}

func (u UserRepository) GetUserByEmail(email string) (models.User, error) {
	return models.User{}, nil
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
