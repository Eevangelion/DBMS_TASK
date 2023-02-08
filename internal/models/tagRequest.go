package models

type TagRequest struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
