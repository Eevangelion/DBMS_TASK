package models

type Joke struct {
	ID           int    `json:"id"`
	Header       string `json:"header"`
	Description  string `json:"description"`
	Rating       int    `json:"rating"`
	AuthorId     int    `json:"author_id"`
	CreationDate string `json:"creation_date"`
}
