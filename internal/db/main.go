package db

import (
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

var UserRepo repositories.IUser = nil
var JokeRepo repositories.IJoke = nil

func init() {
	UserRepo = GetUserRepository()
	JokeRepo = GetJokeRepository()
}
