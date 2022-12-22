package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IJoke interface {
	AddToFavorite(user_id int, joke_id int) (err error)
	DeleteFromFavorite(user_id int, joke_id int) (err error)
	AddTagToJoke(joke_id int, tag_id int) (err error)
	DeleteTagFromJoke(joke_id int, tag_id int) (err error)
	GetUserJokes(user_id int, page int, per_page int, sort_mode string) (jokeOut []models.Joke, err error)
	GetJokeByID(JokeId int) (userOut *models.Joke, err error)
	GetPageOfJokes(page int, per_page int, sort_mode string) (jokes []models.Joke, err error)
	GetJokesByTag(tag_name string) (jokes []models.Joke, err error)
	GetJokesByKeyword(keyword string) (jokes []models.Joke, err error)
	Create(joke *models.Joke) (err error)
	Delete(joke_id int) (err error)
}
