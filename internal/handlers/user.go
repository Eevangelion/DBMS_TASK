package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func UserNameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user *models.User
	username := params["username"]
	user, err := db.UserRepo.GetUserByUsername(username)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
