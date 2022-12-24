package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateJokeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var joke models.Joke
	err := decoder.Decode(&joke)
	if err != nil {
		panic(err)
	}
	id, err := db.JokeRepo.Create(&joke)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DeleteJokeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var joke models.Joke
	err := decoder.Decode(&joke)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.Delete(joke.ID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserJokesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["ID"])
	if err != nil {
		panic(err)
	}
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		panic(err)
	}
	sort_mode := params["Sort"]
	var jokes []models.Joke
	jokes, err = db.JokeRepo.GetUserJokes(user_id, page, pageSize, sort_mode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func GetPageOfJokesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["Page"])
	if err != nil {
		panic(err)
	}
	per_page, err := strconv.Atoi(params["Per_Page"])
	if err != nil {
		panic(err)
	}
	sort_mode := params["Sort"]
	var jokes []models.Joke
	jokes, err = db.JokeRepo.GetPageOfJokes(page, per_page, sort_mode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SearchJokesByTagHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_name := params["Tag"]
	jokes, err := db.JokeRepo.GetJokesByTag(tag_name)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SearchJokesByKeywordHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyword := params["Keyword"]
	jokes, err := db.JokeRepo.GetJokesByKeyword(keyword)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}
