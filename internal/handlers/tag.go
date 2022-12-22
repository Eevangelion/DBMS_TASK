package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tag models.Tag
	err := decoder.Decode(&tag)
	if err != nil {
		panic(err)
	}
	err = db.TagRepo.Create(&tag)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tag models.Tag
	err := decoder.Decode(&tag)
	if err != nil {
		panic(err)
	}
	err = db.TagRepo.Delete(tag.ID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}
