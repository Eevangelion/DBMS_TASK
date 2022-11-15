package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func TagIndexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tag models.Tag
	strId := params["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	tag = db.DB.Tags[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}
