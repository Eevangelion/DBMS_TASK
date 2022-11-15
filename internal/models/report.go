package models

type Report struct {
	ID             int    `json:"id"`
	Description    string `json:"description"`
	ReceiverJokeId int    `json:"receiver_joke_id"`
	SenderId       int    `json:"sender_id"`
	ReceiverId     int    `json:"receiver_id"`
}
