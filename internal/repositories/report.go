package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IReport interface {
	GetReportByID(ReportID int) (reportOut *models.Report, err error)
	Create(report *models.Report) (reportOut *models.Report, err error)
	Delete(report *models.Report) (err error)
}
