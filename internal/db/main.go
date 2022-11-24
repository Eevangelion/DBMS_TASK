package db

import (
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

var UserRepo repositories.IUser = nil

func init() {
	UserRepo = GetUserRepository()
}
