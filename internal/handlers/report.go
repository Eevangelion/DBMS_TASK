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
	setupCors(&w, r)
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
	user, err := db.UserRepo.GetUserByID(report.SenderId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if user.RemainingReports == 0 {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: no reports remains")
		return
	}
	joke, err := db.JokeRepo.GetJokeByID(report.ReceiverJokeId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	err = db.UserRepo.ChangeUserRemainingReports(report.SenderId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
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
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report_id = f["report_id"]
	report, err := db.ReportRepo.GetReportByID(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	if report == nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: no report found")
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
	decoder := json.NewDecoder(r.Body)
	var report_id int
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report_id = f["report_id"]
	report, err := db.ReportRepo.GetReportByID(report_id)
	if report == nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: no report_found")
		return
	}
	err = db.UserRepo.Ban(report.ReceiverId)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	err = db.UserRepo.ChangeUserReportsCount(report.ReceiverId)
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
	decoder := json.NewDecoder(r.Body)
	var report_id int
	var f map[string]int
	err := decoder.Decode(&f)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	report_id = f["report_id"]
	report, err := db.ReportRepo.GetReportByID(report_id)
	if report == nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: no report_found")
		return
	}
	err = db.ReportRepo.Delete(report_id)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
