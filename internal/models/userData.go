package models

type UserData struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Reports          int    `json:"reports"`
	RemainingReports int    `json:"remaining_reports"`
	Role             string `json:"role"`
	UnbanDate        string `json:"unban_date"`
}
