package psql

import (
	"errors"
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type ReportRepository struct {
	report repositories.IReport
}

func (r ReportRepository) GetReportByID(ReportID int) (reportOut *models.Report, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Reports" where id=$1`
	rows, err := DB.Query(qry, ReportID)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get report by ID:", err)
		return nil, err
	}
	var id, receiver_joke_id, sender_id, receiver_id int
	var description string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &description, &receiver_joke_id, &sender_id, &receiver_id)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
	}
	if id != -1 {
		return &models.Report{
			ID:             id,
			Description:    description,
			ReceiverJokeId: receiver_joke_id,
			SenderId:       sender_id,
			ReceiverId:     receiver_id,
		}, nil
	}
	return nil, errors.New("Report with this id does not exist!")
}

func (r ReportRepository) GetAllReports() (reportsOut []models.Report, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Reports"`
	rows, err := DB.Query(qry)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get all reports:", err)
		return nil, err
	}
	for rows.Next() {
		var id, receiver_joke_id, sender_id, receiver_id int
		var description string
		err := rows.Scan(&id, &description, &receiver_joke_id, &sender_id, &receiver_id)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
		NewReport := models.Report{
			ID:             id,
			Description:    description,
			ReceiverJokeId: receiver_joke_id,
			SenderId:       sender_id,
			ReceiverId:     receiver_id,
		}
		reportsOut = append(reportsOut, NewReport)
	}
	return reportsOut, nil
}

func (r ReportRepository) Create(report *models.Report) (reportOut *models.Report, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `INSERT INTO public."Reports" (header, description, rating, author_id) values ($1, $2, $3, $4)`
	result, err := DB.Exec(qry, report.Description, report.ReceiverJokeId, report.SenderId, report.ReceiverId)
	if err != nil {
		log.Println("Error while trying to create report:", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error while trying to create report:", err)
		return nil, err
	}
	reportOut, err = r.GetReportByID(int(id))
	return reportOut, err
}

func (r ReportRepository) Delete(report_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Reports" where id=$1`
	_, err = DB.Exec(qry, report_id)
	if err != nil {
		log.Println("Error while trying to delete report:", err)
		return err
	}
	return nil
}
