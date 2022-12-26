package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IUser interface {
	GetUserByID(user_id int) (userOut *models.User, err error)
	GetUserByUsername(username string) (userOut *models.User, err error)
	GetUserByEmail(Email string) (userOut *models.User, err error)
	GetAll() (users []models.User, err error)
	GetPeopleByKeyword(keyword string, page int, pageSize int) (users []models.User, err error)
	UserChange(user_id int) (err error)

	Ban(user_id int) (err error) // for 1 week (update unban date)
	Create(user *models.User) (id int64, err error)
	Delete(user_id int) (err error)
}
