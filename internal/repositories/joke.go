package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IJoke interface {
	AddToFavorite(user_id int, joke_id int) (err error)
	AddTagToJoke(joke_id int, tag_id int) (err error)
	GetJokeByID(JokeId int) (userOut *models.Joke, err error)
	GetAll() (jokes []models.Joke, err error)
	Create(joke *models.Joke) (jokeOut *models.Joke, err error)
	Delete(joke_id int) (err error)
}
