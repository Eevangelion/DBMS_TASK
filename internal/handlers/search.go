package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/gorilla/mux"

	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	tArg := params["tArg"]
	qArg := params["qArg"]
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
	if tArg == "tag" {
		jokes, err := db.JokeRepo.GetJokesByTag(qArg, page, pageSize)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jokes)
	}
	if tArg == "keyword" {
		jokes, err := db.JokeRepo.GetJokesByKeyword(qArg, page, pageSize)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jokes)
	}
	if tArg == "people" {
		users, err := db.UserRepo.GetPeopleByKeyword(tArg, page, pageSize)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
