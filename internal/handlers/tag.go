package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var tagRequest models.TagRequest
	err := decoder.Decode(&tagRequest)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(tagRequest.UserID)
	if user.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: no permission")
		return
	}
	id, err := db.TagRepo.Create(tagRequest.Name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var tagRequest models.TagRequest
	err := decoder.Decode(&tagRequest)
	log.Println(tagRequest)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(tagRequest.UserID)
	if user.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: no permission")
		return
	}
	err = db.TagRepo.Delete(tagRequest.Name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTagByIDHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag)
}

func GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	tags, err := db.TagRepo.GetAll()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}
