package models

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Reports   int    `json:"reports"`
	Favorites int    `json:"favorites_count"`
	UnbanDate string `json:"unban_date"`
}
