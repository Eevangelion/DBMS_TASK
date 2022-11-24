package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IUser interface {
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)

	Create(user *models.User) (err error)
	Ban(user *models.User) (err error) // for 1 week (update unban date)
	Delete(user *models.User) (err error)
}
