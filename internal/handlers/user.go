package handlers

import (
	"encoding/json"
	"net/http"

	"strconv"
	"time"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	id, err := db.GetUserRepository().Create(&user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err := decoder.Decode(&user_id)
	if err != nil {
		panic(err)
	}
	err = db.GetUserRepository().Delete(user_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		panic(err)
	}
	sortMode := params["sort"]
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err = decoder.Decode(&user_id)
	if err != nil {
		panic(err)
	}
	userOut, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		panic(err)
	}
	favorites, _ := db.JokeRepo.GetUserFavoriteJokes(userOut.ID, 1, 0, "all")
	jokes, _ := db.JokeRepo.GetUserJokes(userOut.ID, page, pageSize, sortMode)
	lastBanDate, err := time.Parse("02.01.2006", userOut.UnbanDate)
	lastBanDate.Add(-1 * time.Hour * 24 * 7)
	userResponse := models.UserResponse{
		ID:          userOut.ID,
		Name:        userOut.Name,
		Reports:     userOut.Reports,
		Favorites:   len(favorites),
		LastBanDate: lastBanDate.Format("02.01.2006"),
		Jokes:       jokes,
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		panic(err)
	}
	keyword_name := params["keyword"]
	jokes, err := db.UserRepo.GetPeopleByKeyWord(keyword_name, page, pageSize)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		panic(err)
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
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(verification)
}
