package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type GitHubOauthToken struct {
	Access_token string
}

func GetGitHubOauthToken(code string) (token string, err error) {
	const rootURl = "https://github.com/login/oauth/access_token"

	conf := config.GetConfig()
	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", conf.GitHubClientID)
	values.Add("client_secret", conf.GitHubClientSecret)

	query := values.Encode()

	queryString := fmt.Sprintf("%s?%s", rootURl, bytes.NewBufferString(query))
	req, err := http.NewRequest("POST", queryString, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("could not retrieve token")
	}

	decoder := json.NewDecoder(res.Body)
	var v map[string]string
	err = decoder.Decode(&v)
	if err != nil {
		return "", err
	}

	token = v["access_token"]

	return token, nil
}

func GetGitHubUser(access_token string) (*models.UserRequestRegisterGithub, error) {
	rootUrl := "https://api.github.com/user"

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	// req.Header.Set("Accept", "application/json")

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	decoder := json.NewDecoder(res.Body)
	var v map[string]interface{}
	err = decoder.Decode(&v)
	if err != nil {
		return nil, err
	}

	user_id := int(v["id"].(float64))

	mail_interface := v["email"]
	mail := ""
	if mail_interface == nil {
		mail = ""
	} else {
		mail = v["email"].(string)
	}

	userBody := &models.UserRequestRegisterGithub{
		ID:                  user_id,
		Name:                "user" + strconv.Itoa(user_id),
		Email:               mail,
		TransformedPassword: access_token,
	}

	return userBody, nil
}
