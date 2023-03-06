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
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var user models.UserRequestRegister
	err := decoder.Decode(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user.Password = utils.GeneratePasswordHash(user.Password)
	id, err := db.GetUserRepository().CreateUser(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	userOut, err := db.GetUserRepository().GetUserByID(int(id))
	jwt_token, refresh_token, err := utils.CreateTokens(userOut)
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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
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
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
	token := r.Header.Get("authorization")
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
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
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
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

func GetUserUnbanDate(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	unban_date, err := db.UserRepo.GetUserUnbanDate(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(unban_date)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	decoder := json.NewDecoder(r.Body)
	var userRequestLogin models.UserRequestLogin
	err := decoder.Decode(&userRequestLogin)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	userRequestLogin.Password = utils.GeneratePasswordHash(userRequestLogin.Password)
	user, err := db.UserRepo.GetUserByUsername(userRequestLogin.Name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user.TransformedPassword != userRequestLogin.Password {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: incorrect password")
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
	userReg := &models.UserRequestRegister{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.TransformedPassword,
	}
	if userOut == nil {
		new_id, err := db.UserRepo.CreateUser(userReg)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		err = db.UserRepo.CreateGithubUserWithID(user.ID, int(new_id))
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		userForOut, err := db.UserRepo.GetUserByID(int(new_id))
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		jwt_token, refresh_token, err := utils.CreateTokens(userForOut)
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
		userForOut, err := db.UserRepo.GetUserByID(userOut.Inner_ID)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
			return
		}
		jwt_token, refresh_token, err := utils.CreateTokens(userForOut)
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
	decoder := json.NewDecoder(r.Body)
	var refresh_token string
	var f map[string]string
	err := decoder.Decode(&f)
	refresh_token = f["refresh_token"]
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = utils.ValidateRefreshToken(refresh_token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id, err := strconv.Atoi(f["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user := models.User{
		ID:   user_id,
		Name: f["name"],
		Role: f["role"],
	}
	jwt_token, new_refresh_token, err := utils.CreateTokens(&user)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	tokens := models.TokensResponse{
		JWTToken:     jwt_token,
		RefreshToken: new_refresh_token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func SubscribeToUserHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	sender_id := claims.User_ID
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	err = decoder.Decode(&receiver_id)
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
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	sender_id := claims.User_ID
	decoder := json.NewDecoder(r.Body)
	var receiver_id int
	err = decoder.Decode(&receiver_id)
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

func GetWhomUserSubscribedTo(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	sender_id := claims.User_ID
	users_id, err := db.UserRepo.GetWhomUserSubscribedTo(sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users_id)
}

func GetWhomUserSubscribedToCount(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	sender_id := claims.User_ID
	amount, err := db.UserRepo.GetWhomUserSubscribedToCount(sender_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(amount)
}

func GetCheckIfUserSubscribed(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	params := mux.Vars(r)
	receiver_id, err := strconv.Atoi(params["receiver_id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	sender_id := claims.User_ID
	check, err := db.UserRepo.GetCheckIfUserSubscribed(sender_id, receiver_id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(check)
}

func ChangeUserNameHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var new_name string
	err = decoder.Decode(&new_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	err = db.UserRepo.ChangeUserName(user_id, new_name)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ChangeUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	token := r.Header.Get("authorization")
	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	user_id := claims.User_ID
	decoder := json.NewDecoder(r.Body)
	var new_transformed_password string
	err = decoder.Decode(&new_transformed_password)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	new_transformed_password = utils.GeneratePasswordHash(new_transformed_password)
	err = db.UserRepo.ChangeUserPassword(user_id, new_transformed_password)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
