package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["username"]
	email := params["email"]
	role := params["role"]
	transformed_password := params["password"]
	user := &models.User{
		ID:                  -14,
		Name:                name,
		Email:               email,
		Reports:             0,
		RemainingReports:    3,
		Role:                role,
		UnbanDate:           "2022-12-22",
		TransformedPassword: transformed_password,
	}
	user, err := db.GetUserRepository().Create(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
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
