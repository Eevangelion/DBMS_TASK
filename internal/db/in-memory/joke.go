package in_memory

import (
	"errors"

	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type JokeRepository struct {
	joke repositories.IJoke
}

var jokeDatabase []models.Joke

func (j JokeRepository) AddToFavorite(id int) (jokeOut *models.Joke, err error) {
	jokeOut, err = j.GetJokeByID(id)
	if err == nil {
		jokeOut.Rating++
	}
	return
}

func (j JokeRepository) GetJokeByID(id int) (jokeOut *models.Joke, err error) {
	for _, v := range jokeDatabase {
		if v.ID == id {
			jokeOut = &v
			return
		}
	}
	err = errors.New("Joke with this ID does not exist!")
	return
}

func (j JokeRepository) GetAll() (jokes []models.Joke, err error) {
	jokes = jokeDatabase
	return
}

func (j JokeRepository) Create(joke *models.Joke) (jokeOut *models.Joke, err error) {
	_, err = j.GetJokeByID(joke.ID)
	if err != nil {
		jokeDatabase = append(jokeDatabase, *joke)
		jokeOut = joke
		err = nil
	} else {
		err = errors.New("Joke with this ID already exists!")
	}
	return
}

func (j JokeRepository) Delete(id int) (jokeOut *models.Joke, err error) {
	_, err = j.GetJokeByID(id)
	if err == nil {
		for i, v := range jokeDatabase {
			if v.ID == id {
				for j := i; j < len(jokeDatabase)-1; j++ {
					jokeDatabase[j] = jokeDatabase[j+1]
				}
				jokeDatabase = jokeDatabase[:len(jokeDatabase)-1]
				jokeOut = &v
				return
			}
		}
	}
	return
}
