package handlers

import (
	"encoding/json"
	"net/http"

	"strconv"
	"time"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	id, err := db.GetUserRepository().Create(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err := decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.GetUserRepository().Delete(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	sortMode := params["sort"]
	if sortMode != "hour" && sortMode != "day" && sortMode != "week" && sortMode != "month" && sortMode != "all" {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	userOut, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	favorites, err := db.JokeRepo.GetUserFavoriteJokes(userOut.ID, 1, 0, "all")
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	jokes, err := db.JokeRepo.GetUserJokes(userOut.ID, page, pageSize, sortMode)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	lastBanDate, err := time.Parse("02.01.2006", userOut.UnbanDate)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	lastBanDate = lastBanDate.Add(-1 * time.Hour * 24 * 7)
	userResponse := models.UserResponse{
		ID:          userOut.ID,
		Name:        userOut.Name,
		Reports:     userOut.Reports,
		Favorites:   len(favorites),
		LastBanDate: lastBanDate.Format("02.01.2006"),
		Jokes:       jokes,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	keyword := params["keyword"]
	jokes, err := db.UserRepo.GetPeopleByKeyWord(keyword, page, pageSize)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jokes)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var password string
	err = decoder.Decode(&password)
	user, err := db.UserRepo.GetUserByID(userID)
	verification := true
	if user.TransformedPassword == password {
		verification = true
	} else {
		verification = false
	}
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verification)
}
