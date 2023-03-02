package handlers

import (
	"encoding/json"
	"log"
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
	var user models.User
	err := decoder.Decode(&user)
	log.Println(user)
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

func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
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
	setupCors(&w)
	params := mux.Vars(r)
	str_id := params["id"]
	id, err := strconv.Atoi(str_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(id)
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

func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
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
	setupCors(&w)
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
	setupCors(&w)
	params := mux.Vars(r)
	code := params["code"]
	token, err := utils.GetGitHubOauthToken(code)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	user, err := utils.GetGitHubUser(token)
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
		jwt_token, refresh_token, err := utils.CreateTokens(user)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error while creating token")
			return
		}
		tokens := models.TokensResponse{
			JWTToken:     jwt_token,
			RefreshToken: refresh_token,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tokens)
	} else {
		user, err = db.UserRepo.GetUserByID(userOut.Inner_ID)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		jwt_token, refresh_token, err := utils.CreateTokens(user)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error while creating token")
			return
		}
		tokens := models.TokensResponse{
			JWTToken:     jwt_token,
			RefreshToken: refresh_token,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tokens)
	}
}

func RefreshUser(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("Authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var refreshToken string
	err = decoder.Decode(&refreshToken)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user := models.User{
		ID:   claims.User_ID,
		Name: claims.UserName,
		Role: claims.Role,
	}
	jwt_token, refresh_token, err := utils.CreateTokens(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	tokens := models.TokensResponse{
		JWTToken:     jwt_token,
		RefreshToken: refresh_token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func SubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	receiver_id = f["receiver_id"]
	sender_id = f["sender_id"]
	err = db.JokeRepo.SubscribeToUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnSubscribeFromUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	var sender_id int
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	receiver_id = f["receiver_id"]
	sender_id = f["sender_id"]
	err = db.JokeRepo.UnSubscribeFromUser(receiver_id, sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ChangeUserNameHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var new_name string
	var f map[string]string
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user_id, err := strconv.Atoi(f["user_id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	new_name = f["name"]
	err = db.UserRepo.ChangeUserName(user_id, new_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ChangeUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var f map[string]string
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user_id, err := strconv.Atoi(f["user_id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	new_transformed_password := f["transformed_password"]
	err = db.UserRepo.ChangeUserPassword(user_id, new_transformed_password)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
