package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"

	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
)

func CreateJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var joke models.Joke
	err := decoder.Decode(&joke)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	id, err := db.JokeRepo.Create(&joke)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func DeleteJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var joke_id int
	var user_id int
	err := decoder.Decode(&joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	joke, err := db.JokeRepo.GetJokeByID(joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user.Role != "admin" && joke.AuthorId != user_id {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.Delete(joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	username := params["username"]
	pageURL := r.URL.Query().Get("pageArg")
	pageSizeURL := r.URL.Query().Get("pageSize")
	var page, pageSize int
	if pageURL == "" {
		page = 1
		pageSize = 5
	} else {
		var err error
		page, err = strconv.Atoi(pageURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
		pageSize, err = strconv.Atoi(pageSizeURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
	}
	sortMode := r.URL.Query().Get("sortMode")
	if sortMode == "" {
		sortMode = "new"
	} else {
		if sortMode != "new" && sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "alltime" {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Wrong feed query parameter")
			return
		}
	}
	user, err := db.UserRepo.GetUserByUsername(username)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	jokes, amount, err := db.JokeRepo.GetUserJokes(user.ID, page, pageSize, sortMode)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.JokeResponse{
		Jokes:  jokes,
		Amount: amount,
	})
}

func GetPageOfJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	pageURL := r.URL.Query().Get("pageArg")
	pageSizeURL := r.URL.Query().Get("pageSize")
	var page, pageSize int
	if pageURL == "" {
		page = 1
		pageSize = 5
	} else {
		var err error
		page, err = strconv.Atoi(pageURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
		pageSize, err = strconv.Atoi(pageSizeURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
	}
	sortMode := r.URL.Query().Get("sortMode")
	if sortMode == "" {
		sortMode = "new"
	} else {
		if sortMode != "new" && sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "alltime" {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Wrong feed query parameter")
			return
		}
	}
	jokes, amount, err := db.JokeRepo.GetPageOfJokes(page, pageSize, sortMode)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.JokeResponse{
		Jokes:  jokes,
		Amount: amount,
	})
}

func GetJokeTagsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	tags, err := db.JokeRepo.GetJokeTags(joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}

func AddToFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.AddToFavorite(user_id, joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteFromFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.DeleteFromFavorite(user_id, joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserFavoriteJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	pageURL := r.URL.Query().Get("pageArg")
	pageSizeURL := r.URL.Query().Get("pageSize")
	var page, pageSize int
	if pageURL == "" {
		page = 1
		pageSize = 5
	} else {
		var err error
		page, err = strconv.Atoi(pageURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
		pageSize, err = strconv.Atoi(pageSizeURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
	}
	sortMode := r.URL.Query().Get("sortMode")
	if sortMode == "" {
		sortMode = "new"
	} else {
		if sortMode != "new" && sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "alltime" {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Wrong feed query parameter")
			return
		}
	}
	jokes, amount, err := db.JokeRepo.GetUserFavoriteJokes(user_id, page, pageSize, sortMode)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.JokeResponse{
		Jokes:  jokes,
		Amount: amount,
	})
}

func SubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.SubscribeToUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnSubscribeFromUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.UnSubscribeFromUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserSubscribedJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	pageURL := r.URL.Query().Get("pageArg")
	pageSizeURL := r.URL.Query().Get("pageSize")
	var page, pageSize int
	if pageURL == "" {
		page = 1
		pageSize = 5
	} else {
		var err error
		page, err = strconv.Atoi(pageURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
		pageSize, err = strconv.Atoi(pageSizeURL)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			return
		}
	}
	sortMode := r.URL.Query().Get("sortMode")
	if sortMode == "" {
		sortMode = "new"
	} else {
		if sortMode != "new" && sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "alltime" {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Wrong feed query parameter")
			return
		}
	}
	jokes, amount, err := db.JokeRepo.GetUserSubribedJokes(receiver_id, page, pageSize, sortMode)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.JokeResponse{
		Jokes:  jokes,
		Amount: amount,
	})
}

func AddTagToJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var tag_id int
	err = decoder.Decode(&tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.AddTagToJoke(joke_id, tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteTagToJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var tag_id int
	err = decoder.Decode(&tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.DeleteTagFromJoke(joke_id, tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	joke_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	joke, err := db.JokeRepo.GetJokeByID(joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(joke)
}
