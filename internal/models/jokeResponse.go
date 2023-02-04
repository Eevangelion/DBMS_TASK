package models

type JokeResponse struct {
	Jokes  []Joke `json:"jokes"`
	Amount int    `json:"amount"`
}
