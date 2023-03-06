package models

type ReportRequest struct {
	Description    string `json:"description"`
	ReceiverJokeId int    `json:"receiver_joke_id"`
	SenderId       int    `json:"sender_id"`
}
