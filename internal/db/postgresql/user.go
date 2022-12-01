package psql

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type UserRepository struct {
	user repositories.IUser
}

func (u UserRepository) GetUserByUsername(username string) (userOut *models.User, err error) {
	return
}

func (u UserRepository) GetUserByEmail(email string) (userOut *models.User, err error) {
	return
}

func (u UserRepository) GetUserByID(id int) (userOut *models.User, err error) {
	return
}

func (u UserRepository) GetAll() (users []models.User, err error) {
	return
}

func (u UserRepository) Create(user *models.User) (userOut *models.User, err error) {
	return
}

func (u UserRepository) Ban(id int) (userOut *models.User, err error) {
	return
}

func (u UserRepository) Delete(id int) (userOut *models.User, err error) {
	return
}
