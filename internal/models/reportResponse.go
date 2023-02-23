package models

type ReportResponse struct {
	Reports []Report `json:"reports"`
	Amount  int      `json:"amount"`
}
