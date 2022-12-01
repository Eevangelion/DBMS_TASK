package psql

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type JokeRepository struct {
	joke repositories.IJoke
}

func (j JokeRepository) AddToFavorite(id int) (jokeOut *models.Joke, err error) {
	return
}

func (j JokeRepository) GetJokeByID(id int) (jokeOut *models.Joke, err error) {
	return
}

func (j JokeRepository) GetAll() (jokes []models.Joke, err error) {
	return
}

func (j JokeRepository) Create(joke *models.Joke) (jokeOut *models.Joke, err error) {
	return
}

func (j JokeRepository) Delete(id int) (jokeOut *models.Joke, err error) {
	return
}
