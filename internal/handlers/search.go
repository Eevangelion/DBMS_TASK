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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	params := mux.Vars(r)
	tArg := params["t"]
	qArg := params["q"]
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
	if tArg == "tag" {
		jokes, amount, err := db.JokeRepo.GetJokesByTag(qArg, page, pageSize, sortMode)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.JokeResponse{
			Jokes:  jokes,
			Amount: amount,
		})
		return
	}
	if tArg == "keyword" {
		jokes, amount, err := db.JokeRepo.GetJokesByKeyword(qArg, page, pageSize, sortMode)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.JokeResponse{
			Jokes:  jokes,
			Amount: amount,
		})
		return
	}
	if tArg == "people" {
		users, _, err := db.UserRepo.GetPeopleByKeyword(qArg, page, pageSize)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
		return
	} else {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: bad tArg")
		return
	}
}
