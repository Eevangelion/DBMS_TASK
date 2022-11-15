package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakagam1/DBMS_TASK/internal/db"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/gorilla/mux"
)

func ReportIndexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var report models.Report
	strId := params["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	report = db.DB.Reports[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
