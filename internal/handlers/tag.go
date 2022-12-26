package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tag models.Tag
	err := decoder.Decode(&tag)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	id, err := db.TagRepo.Create(tag.Name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_id, err := strconv.Atoi(params["tagID"])
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	err = db.TagRepo.Delete(tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTagByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_id, err := strconv.Atoi(params["tagID"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	tag, err := db.TagRepo.GetTagByID(tag_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag)
}

func GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := db.TagRepo.GetAll()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}
