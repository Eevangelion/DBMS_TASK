package handlers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/utils"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userRequest models.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	var user models.User
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.TransformedPassword = userRequest.TransformedPassword
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

func GetUserDataByNameHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	username := params["username"]
	user, err := db.UserRepo.GetUserByUsername(username)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user == nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: no user found")
		return
	}
	userData := models.UserData{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		Role:             user.Role,
		Reports:          user.Reports,
		RemainingReports: user.RemainingReports,
		UnbanDate:        user.UnbanDate,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}

func GetUserDataByIDHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	user_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	userOut, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	_, amount, err := db.JokeRepo.GetUserFavoriteJokes(userOut.ID, 1, 0, "new")
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	userResponse := models.UserResponse{
		ID:        userOut.ID,
		Name:      userOut.Name,
		Role:      userOut.Role,
		Reports:   userOut.Reports,
		Favorites: amount,
		UnbanDate: userOut.UnbanDate,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verification)
}

func GetGithubUser(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var code string
	err := decoder.Decode(&code)
	token, err := utils.GetGitHubOauthToken(code)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	user, err := utils.GetGitHubUser(token.Access_token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	userOut, err := db.UserRepo.GetUserByGithubID(user.ID)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if userOut == nil {
		new_id, err := db.UserRepo.Create(user)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		err = db.UserRepo.CreateGithubUserWithID(user.ID, int(new_id))
		userResponse := models.UserResponse{
			ID:        int(new_id),
			Name:      user.Name,
			Role:      user.Role,
			Reports:   user.Reports,
			Favorites: 0,
			UnbanDate: user.UnbanDate,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userResponse)
		json.NewEncoder(w).Encode(token)
	} else {
		_, amount, err := db.JokeRepo.GetUserFavoriteJokes(user.ID, 1, 0, "new")
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		userResponse := models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Role:      user.Role,
			Reports:   user.Reports,
			Favorites: amount,
			UnbanDate: user.UnbanDate,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userResponse)
		json.NewEncoder(w).Encode(token)
	}
}

func SubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.SubscribeToUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnSubscribeFromUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	err := decoder.Decode(&receiver_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = decoder.Decode(&sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.JokeRepo.UnSubscribeFromUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
