package models

type JokeRequest struct {
	Header      string `json:"header"`
	Description string `json:"description"`
	AuthorId    int    `json:"author_id"`
	Tags        []Tag  `json:"tags"`
}
