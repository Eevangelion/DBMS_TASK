package models

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Reports     int    `json:"reports"`
	Favorites   int    `json:"remainingReports"`
	LastBanDate string `json:"lastBanDate"`
}
