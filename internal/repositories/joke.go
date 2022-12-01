package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IJoke interface {
	AddToFavorite(id int) (jokeOut *models.Joke, err error)

	Create(joke *models.Joke) (jokeOut *models.Joke, err error)
	GetAll() (jokes []models.Joke, err error)
	Delete(id int) (jokeOut *models.Joke, err error)
}
