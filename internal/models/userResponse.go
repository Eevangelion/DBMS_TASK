package models

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Reports     int    `json:"reports"`
	Favorites   int    `json:"remaining_reports"`
	LastBanDate string `json:"unban_date"`
	Jokes       []Joke `json:"jokes"`
}
