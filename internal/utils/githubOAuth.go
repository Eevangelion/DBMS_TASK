package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func GetGitHubOauthToken(code string) (*GitHubOauthToken, error) {
	const rootURl = "https://github.com/login/oauth/access_token"
	config := config.GetConfig()

	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", config.GitHubClientID)
	values.Add("client_secret", config.GitHubClientSecret)

	query := values.Encode()

	queryString := fmt.Sprintf("%s?%s", rootURl, bytes.NewBufferString(query))
	req, err := http.NewRequest("POST", queryString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	parsedQuery, err := url.ParseQuery(string(resBody))
	if err != nil {
		return nil, err
	}

	tokenBody := &GitHubOauthToken{
		Access_token: parsedQuery["access_token"][0],
	}

	return tokenBody, nil
}

func GetGitHubUser(access_token string) (*models.User, error) {
	rootUrl := "https://api.github.com/user"

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

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

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var GitHubUserRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GitHubUserRes); err != nil {
		return nil, err
	}

	user_id, err := strconv.Atoi(GitHubUserRes["id"].(string))

	userBody := &models.User{
		ID:                  user_id,
		Name:                GitHubUserRes["login"].(string),
		Email:               GitHubUserRes["email"].(string),
		Reports:             0,
		RemainingReports:    0,
		UnbanDate:           "date",
		TransformedPassword: access_token,
	}

	return userBody, nil
}