package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tag models.Tag
	err := decoder.Decode(&tag)
	if err != nil {
		panic(err)
	}
	id, err := db.TagRepo.Create(tag.Name)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_id, err := strconv.Atoi(params["tagID"])
	if err != nil {
		panic(err)
	}
	err = db.TagRepo.Delete(tag_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetTagByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_id, err := strconv.Atoi(params["tagID"])
	if err != nil {
		panic(err)
	}
	tag, err := db.TagRepo.GetTagByID(tag_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

func GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := db.TagRepo.GetAll()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
