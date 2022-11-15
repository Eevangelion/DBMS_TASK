package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func JokeIndexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var joke models.Joke
	strId := params["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	joke = db.DB.Jokes[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joke)
}
