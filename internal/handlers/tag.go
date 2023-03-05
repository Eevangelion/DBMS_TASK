package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
	"github.com/Sakagam1/DBMS_TASK/internal/utils"
	"github.com/gorilla/mux"
)

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	if claims.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var tag_name string
	err = decoder.Decode(&tag_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	id, err := db.TagRepo.Create(tag_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	if claims.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var tag_name string
	err = decoder.Decode(&tag_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.TagRepo.Delete(tag_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTagByIDHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	tags, err := db.TagRepo.GetAllTags()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}
