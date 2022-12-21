package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

func CreateReportHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var report models.Report
	err := decoder.Decode(&report)
	if err != nil {
		panic(err)
	}
	var reportOut *models.Report
	reportOut, err = db.ReportRepo.Create(&report)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reportOut)
}

func DeleteReportHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var report models.Report
	err := decoder.Decode(&report)
	if err != nil {
		panic(err)
	}
	var reportOut *models.Report
	err = db.ReportRepo.Delete(&report)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reportOut)
}

func GetReportHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var report models.Report
	err := decoder.Decode(&report)
	if err != nil {
		panic(err)
	}
	var reportOut *models.Report
	reportOut, err = db.ReportRepo.GetReportByID(report.ID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reportOut)
}

func GetAllReportsHandler(w http.ResponseWriter, r *http.Request) {
	reportOut, err := db.ReportRepo.GetAllReports()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reportOut)
}
