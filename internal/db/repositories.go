package db

import (
	"github.com/Sakagam1/DBMS_TASK/internal/config"
	// in_memory "github.com/Sakagam1/DBMS_TASK/internal/db/in-memory"
	psql "github.com/Sakagam1/DBMS_TASK/internal/db/postgresql"
	repositories "github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

func GetUserRepository() repositories.IUser {
	if UserRepo == nil {
		if config.GetConfig().DebugMode == true {
			UserRepo = nil
			//UserRepo = new(in_memory.UserRepository)
		} else {
			UserRepo = new(psql.UserRepository)
		}
	}
	return UserRepo
}
func GetJokeRepository() repositories.IJoke {
	if JokeRepo == nil {
		if config.GetConfig().DebugMode == true {
			JokeRepo = nil
			//JokeRepo = new(in_memory.UserRepository)
		} else {
			JokeRepo = new(psql.JokeRepository)
		}
	}
	return JokeRepo
}

func GetTagRepository() repositories.ITag {
	if TagRepo == nil {
		if config.GetConfig().DebugMode == true {
			TagRepo = nil
			//TagRepo = new(in_memory.UserRepository)
		} else {
			TagRepo = new(psql.TagRepository)
		}
	}
	return TagRepo
}

func GetReportRepository() repositories.IReport {
	if ReportRepo == nil {
		if config.GetConfig().DebugMode == true {
			ReportRepo = nil
			//ReportRepo = new(in_memory.UserRepository)
		} else {
			ReportRepo = new(psql.ReportRepository)
		}
	}
	return ReportRepo
}
