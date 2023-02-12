package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IUser interface {
	GetUserByID(user_id int) (userOut *models.User, err error)
	GetUserByUsername(username string) (userOut *models.User, err error)
	GetUserByEmail(Email string) (userOut *models.User, err error)
	GetAll() (users []models.User, err error)
	GetPeopleByKeyword(keyword string, page int, pageSize int) (users []models.UserResponseSearch, err error)
	ChangeUserReport(user_id int) (err error)
	ChangeUserName(user_id int, new_name string) (err error)
	ChangeUserPassword(user_id int, new_transformed_password string) (err error)
	GetUserByGithubID(user_id int) (userOut *models.GitUser, err error)
	CreateGithubUserWithID(user_id int, inner_id int) (err error)
	GetSubscribedPeopleCount(user_id int) (amount int, err error)
	GetUserJokesCount(user_id int) (amount int, err error)

	Ban(user_id int) (err error) // for 1 week (update unban date)
	Create(user *models.User) (id int64, err error)
	Delete(user_id int) (err error)
}
