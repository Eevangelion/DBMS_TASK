package db

import (
	"github.com/Sakagam1/DBMS_TASK/internal/config"
	in_memory "github.com/Sakagam1/DBMS_TASK/internal/db/in-memory"
	psql "github.com/Sakagam1/DBMS_TASK/internal/db/postgresql"
	repositories "github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

func GetUserRepository() repositories.IUser {
	if UserRepo == nil {
		UserRepo = new(in_memory.UserRepository)
		if config.GetConfig().DebugMode == true {
			UserRepo = new(in_memory.UserRepository)
		} else {
			UserRepo = new(psql.UserRepository)
		}
	}
	return UserRepo
}
