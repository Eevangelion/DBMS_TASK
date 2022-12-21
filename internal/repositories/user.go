package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IUser interface {
	GetUserByID(user_id int) (userOut *models.User, err error)
	GetUserByUsername(username string) (userOut *models.User, err error)
	GetUserByEmail(Email string) (userOut *models.User, err error)
	Create(user *models.User) (userOut *models.User, err error)
	GetAll() (users []models.User, err error)
	GetPeopleByKeyWord(keyword string) (users []models.User, err error)

	Ban(user *models.User) (err error) // for 1 week (update unban date)
	Delete(user *models.User) (err error)
}
