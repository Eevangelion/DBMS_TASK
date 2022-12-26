package models

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Reports   int    `json:"reports"`
	Favorites int    `json:"remaining_reports"`
	UnbanDate string `json:"unban_date"`
}
