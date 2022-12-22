package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	err = db.GetUserRepository().Create(&user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	var userOut *models.User
	userOut, err = db.UserRepo.GetUserByID(user.ID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(userOut)
}

func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyword_name := params["KeyWord"]
	var jokes []models.User
	jokes, err := db.UserRepo.GetPeopleByKeyWord(keyword_name)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}
