package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"

	customHTTP "github.com/Sakagam1/DBMS_TASK/internal/http"
)

func CreateReportHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reportRequest models.ReportRequest
	err := decoder.Decode(&reportRequest)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	var report models.Report
	report.Description = reportRequest.Description
	report.ReceiverJokeId = reportRequest.ReceiverJokeId
	report.SenderId = reportRequest.SenderId
	joke, err := db.JokeRepo.GetJokeByID(report.ReceiverJokeId)
	report.ReceiverId = joke.AuthorId
	id, err := db.ReportRepo.Create(&report)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteReportHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var report_id int
	var user_id int
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report_id = f["report_id"]
	user_id = f["user_id"]
	report, err := db.ReportRepo.GetReportByID(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if report == nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: no report found")
		return
	}
	user, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user == nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: no user found")
		return
	}
	if user.Role != "admin" && user_id != report.SenderId {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	err = db.ReportRepo.Delete(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetReportByIDHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	report_id, err := strconv.Atoi(params["id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	reportOut, err := db.ReportRepo.GetReportByID(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reportOut)
}

func GetAllReportsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	decoder := json.NewDecoder(r.Body)
	var user_id int
	err := decoder.Decode(&user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	user, err := db.UserRepo.GetUserByID(user_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user.Role != "admin" {
		customHTTP.NewErrorResponse(w, http.StatusForbidden, "Error: "+err.Error())
		return
	}
	reportOut, err := db.ReportRepo.GetAllReports()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reportOut)
}

func ApplyReportHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	report_id, err := strconv.Atoi(params["report_id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report, err := db.ReportRepo.GetReportByID(report_id)
	err = db.UserRepo.Ban(report.ReceiverId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	err = db.UserRepo.UserChange(report.SenderId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	err = db.ReportRepo.Delete(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DenyReportHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	report_id, err := strconv.Atoi(params["report_id"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report, err := db.ReportRepo.GetReportByID(report_id)
	err = db.UserRepo.UserChange(report.SenderId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	err = db.ReportRepo.Delete(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
