package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IReport interface {
	GetReportByID(ReportID int) (reportOut *models.Report, err error)
	GetAllReports() (reportsOut *models.ReportResponse, err error)
	Create(report *models.Report) (id int64, err error)
	Delete(report_id int) (err error)
}
