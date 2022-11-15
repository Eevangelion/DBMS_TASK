package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	strId := params["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	user = db.DB.Users[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
