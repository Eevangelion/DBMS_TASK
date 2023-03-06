package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/utils"
	"github.com/gorilla/mux"

	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
)

func CreateJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	if _, err := utils.ValidateAccessToken(token); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var jokeRequest models.JokeRequest
	err := decoder.Decode(&jokeRequest)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	var joke models.Joke
	joke.AuthorId = jokeRequest.AuthorId
	joke.Description = jokeRequest.Description
	joke.Header = jokeRequest.Header
	var tags []models.Tag
	tags = jokeRequest.Tags
	id, err := db.JokeRepo.Create(&joke)
	for _, tag := range tags {
		err = db.JokeRepo.AddTagToJoke(int(id), tag.ID)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
	}
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	for _, tag := range jokeRequest.Tags {
		err = db.JokeRepo.AddTagToJoke(int(id), tag.ID)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func DeleteJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	if _, err := utils.ValidateAccessToken(token); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var joke_id int
	err := decoder.Decode(&joke_id)
	joke, err := db.JokeRepo.GetJokeByID(joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if joke == nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: no joke found")
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	if _, err := utils.ValidateAccessToken(token); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	params := mux.Vars(r)
	username := params["username"]
	pageURL := r.URL.Query().Get("page")
	pageSizeURL := r.URL.Query().Get("page_size")
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
		if pageSizeURL == "" {
			pageSize = 5
		} else {
			pageSize, err = strconv.Atoi(pageSizeURL)
			if err != nil {
				customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
				return
			}
		}
	}
	sortMode := r.URL.Query().Get("sort")
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
	if user == nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: no user found")
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	if _, err := utils.ValidateAccessToken(token); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	pageURL := r.URL.Query().Get("page")
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
		if pageSizeURL == "" {
			pageSize = 5
		} else {
			pageSize, err = strconv.Atoi(pageSizeURL)
			if err != nil {
				customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
				return
			}
		}
	}
	sortMode := r.URL.Query().Get("sort")
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	if _, err := utils.ValidateAccessToken(token); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var joke_id int
	user_id := claims.User_ID
	err = decoder.Decode(&joke_id)
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	decoder := json.NewDecoder(r.Body)
	var joke_id int
	err = decoder.Decode(&joke_id)
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

func CheckIfInFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	var f map[string]int
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	var joke_id int
	joke_id = f["joke_id"]
	amount, err := db.JokeRepo.CheckIfInFavorite(user_id, joke_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	json.NewEncoder(w).Encode(amount != 0)
	w.WriteHeader(http.StatusOK)
}

func GetUserFavoriteJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
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

func GetUserSubscribedJokesHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	pageURL := r.URL.Query().Get("page")
	pageSizeURL := r.URL.Query().Get("page_size")
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
		if pageSizeURL == "" {
			pageSize = 5
		} else {
			pageSize, err = strconv.Atoi(pageSizeURL)
			if err != nil {
				customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
				return
			}
		}
	}
	sortMode := r.URL.Query().Get("sort")
	if sortMode == "" {
		sortMode = "new"
	} else {
		if sortMode != "new" && sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "alltime" {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Wrong feed query parameter")
			return
		}
	}
	jokes, amount, err := db.JokeRepo.GetUserSubscribedJokes(user_id, page, pageSize, sortMode)
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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

func DeleteTagFromJokeHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
