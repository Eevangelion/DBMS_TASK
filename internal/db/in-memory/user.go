package in_memory

import (
	"errors"

	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type UserRepository struct {
	user repositories.IUser
}

var userDatabase []models.User

func (u UserRepository) GetUserByUsername(username string) (userOut *models.User, err error) {
	for _, v := range userDatabase {
		if v.Name == username {
			userOut = &v
			return
		}
	}
	err = errors.New("User with this username does not exist!")
	return
}

func (u UserRepository) GetUserByEmail(email string) (userOut *models.User, err error) {
	for _, v := range userDatabase {
		if v.Email == email {
			userOut = &v
			return
		}
	}
	err = errors.New("User with this email does not exist!")
	return
}

func (u UserRepository) GetUserByID(id int) (userOut *models.User, err error) {
	for _, v := range userDatabase {
		if v.ID == id {
			userOut = &v
			return
		}
	}
	err = errors.New("User with this ID does not exist!")
	return
}

func (u UserRepository) GetAll() (users []models.User, err error) {
	users = userDatabase
	return
}

func (u UserRepository) Create(user *models.User) (userOut *models.User, err error) {
	_, err = u.GetUserByID(user.ID)
	if err != nil {
		userDatabase = append(userDatabase, *user)
		userOut = user
		err = nil
	} else {
		err = errors.New("User with this ID already exists!")
	}
	return
}

func (u UserRepository) Ban(id int) (userOut *models.User, err error) {
	return
}

func (u UserRepository) Delete(id int) (userOut *models.User, err error) {
	_, err = u.GetUserByID(id)
	if err == nil {
		for i, v := range userDatabase {
			if v.ID == id {
				for j := i; j < len(userDatabase)-1; j++ {
					userDatabase[j] = userDatabase[j+1]
				}
				userDatabase = userDatabase[:len(userDatabase)-1]
				userOut = &v
				break
			}
		}
	}
	return
}
