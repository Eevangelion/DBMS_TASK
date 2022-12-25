package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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
	id, err := db.ReportRepo.Create(&report)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DeleteReportHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	report_id, err := strconv.Atoi(params["reportID"])
	if err != nil {
		panic(err)
	}
	err = db.ReportRepo.Delete(report_id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func GetReportByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	report_id, err := strconv.Atoi(params["reportID"])
	if err != nil {
		panic(err)
	}
	reportOut, err := db.ReportRepo.GetReportByID(report_id)
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
