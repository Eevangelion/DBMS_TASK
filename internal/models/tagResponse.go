package models

type TagResponse struct {
	Tags   []Tag `json:"tags"`
	Amount int   `json:"amount"`
}
