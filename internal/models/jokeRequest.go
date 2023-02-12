package models

type JokeRequest struct {
	Tags        []Tag  `json:"tags"`
	Header      string `json:"header"`
	Description string `json:"description"`
	AuthorId    int    `json:"author_id"`
}
