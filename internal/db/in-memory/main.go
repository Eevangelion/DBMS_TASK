package in_memory

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type DataBase struct {
	Jokes   map[int]models.Joke
	Users   map[int]models.User
	Tags    map[int]models.Tag
	Reports map[int]models.Report
}

func SetupDB() *DataBase {
	var db *DataBase = &DataBase{
		Jokes: map[int]models.Joke{
			1: {
				ID:          1,
				Header:      "Смешной заголовок.",
				Description: "Смешная шутка.",
				Rating:      5,
				AuthorId:    1,
			},
			2: {
				ID:          2,
				Header:      "Ещё один смешной заголовок.",
				Description: "Ещё одна смешная шутка.",
				Rating:      4,
				AuthorId:    3,
			},
		},
		Users: map[int]models.User{
			1: {
				ID:                  1,
				Name:                "Ivan",
				Email:               "ivan@mail.com",
				Reports:             2,
				RemainingReports:    1,
				Role:                "Пользователь",
				UnbanDate:           "01.01.1970",
				TransformedPassword: "ASkodasi213mdasok",
			},
			2: {
				ID:                  2,
				Name:                "Alex",
				Email:               "alex@mail.com",
				Reports:             3,
				RemainingReports:    0,
				Role:                "Пользователь",
				UnbanDate:           "31.12.2022",
				TransformedPassword: "kodsADk213kl4SDakl",
			},
			3: {
				ID:                  3,
				Name:                "Max",
				Email:               "max@mail.com",
				Reports:             0,
				RemainingReports:    3,
				Role:                "Администратор",
				UnbanDate:           "01.01.1970",
				TransformedPassword: "3SJRijsmvm23kdsoj4",
			},
		},
		Tags: map[int]models.Tag{
			1: {
				ID:   1,
				Name: "Очень смешная шутка",
			},
			2: {
				ID:   2,
				Name: "Не смешная шутка",
			},
			3: {
				ID:   3,
				Name: "Шутка про Штирлица",
			},
		},
		Reports: map[int]models.Report{},
	}
	return db
}
