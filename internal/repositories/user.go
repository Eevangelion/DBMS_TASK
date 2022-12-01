package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IUser interface {
	GetUserByUsername(username string) (userOut *models.User, err error)
	GetUserByEmail(email string) (userOut *models.User, err error)
	GetUserByID(id int) (userOut *models.User, err error)
	GetAll() (users []models.User, err error)

	Create(user *models.User) (userOut *models.User, err error)
	Ban(id int) (userOut *models.User, err error) // for 1 week (update unban date)
	Delete(id int) (userOut *models.User, err error)
}
