package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

func CreateJokeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var joke models.Joke
	err := decoder.Decode(&joke)
	if err != nil {
		panic(err)
	}
	var jokeOut *models.Joke
	jokeOut, err = db.JokeRepo.Create(&joke)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(jokeOut)
}

func DeleteJokeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var joke models.Joke
	err := decoder.Decode(&joke)
	if err != nil {
		panic(err)
	}
	var jokeOut *models.Joke
	jokeOut, err = db.JokeRepo.Delete(joke.ID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(jokeOut)
}

func GetJokes(w http.ResponseWriter, r *http.Request) {

	var jokes []models.Joke

	jokes, err := db.JokeRepo.GetAll()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}
