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
	var joke_id int
	err := decoder.Decode(&joke_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.Delete(joke_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserJokesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["userID"])
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
	sortMode := params["sort"]
	var jokes []models.Joke
	jokes, err = db.JokeRepo.GetUserJokes(user_id, page, pageSize, sortMode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func GetPageOfJokesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		panic(err)
	}
	sortMode := params["sort"]
	var jokes []models.Joke
	jokes, err = db.JokeRepo.GetPageOfJokes(page, pageSize, sortMode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SearchJokesByTagHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_name := params["tag"]
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	sortMode := params["sort"]
	if err != nil {
		panic(err)
	}
	jokes, err := db.JokeRepo.GetJokesByTag(tag_name, page, pageSize, sortMode)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SearchJokesByKeywordHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyword := params["keyword"]
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		panic(err)
	}
	sortMode := params["sort"]
	jokes, err := db.JokeRepo.GetJokesByKeyword(keyword, page, pageSize, sortMode)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func AddToFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["jokeID"])
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.AddToFavorite(user_id, joke_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func DeleteFromFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["jokeID"])
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.DeleteFromFavorite(user_id, joke_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserFavoriteJokesHandler(w http.ResponseWriter, r *http.Request) {
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
	sortMode := params["sort"]
	var jokes []models.Joke
	jokes, err = db.JokeRepo.GetUserFavoriteJokes(user_id, page, pageSize, sortMode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.SubscribeToUser(receiver_id, sender_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func UnSubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.UnSubscribeFromUser(receiver_id, sender_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserSubscribedJokesHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := mux.Vars(r)
	var receiver_id int
	err := decoder.Decode(&receiver_id)
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
	sortMode := params["sort"]
	jokes, err := db.JokeRepo.GetUserSubribedJokes(receiver_id, page, pageSize, sortMode)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func AddTagToJokeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["jokeID"])
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(r.Body)
	var tag_id int
	err = decoder.Decode(&tag_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.AddTagToJoke(joke_id, tag_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func DeleteTagToJokeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["jokeID"])
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(r.Body)
	var tag_id int
	err = decoder.Decode(&tag_id)
	if err != nil {
		panic(err)
	}
	err = db.JokeRepo.DeleteTagFromJoke(joke_id, tag_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["jokeID"])
	if err != nil {
		panic(err)
	}
	joke, err := db.JokeRepo.GetJokeByID(joke_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joke)
}
