package psql

import (
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type ReportRepository struct {
	report repositories.IReport
}

func (r ReportRepository) GetReportByID(report_id int) (reportOut *models.Report, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	var id, receiver_joke_id, sender_id, receiver_id int
	var description string
	qry := `select * from public."Reports" where id=$1`
	qry2 := `select count("Reports".id) from public."Reports" where id=$1`
	err = DB.QueryRow(qry2, report_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get report by ID (amount):", err)
	}
	if amount == 0 {
		return reportOut, nil
	}
	err = DB.QueryRow(qry, report_id).Scan(&id, &description, &receiver_joke_id, &sender_id, &receiver_id)
	if err != nil {
		log.Println("Error while trying to get report by ID:", err)
		return nil, err
	}
	return &models.Report{
		ID:             report_id,
		Description:    description,
		ReceiverJokeId: receiver_joke_id,
		SenderId:       sender_id,
		ReceiverId:     receiver_id,
	}, nil
}

func (r ReportRepository) GetAllReports() (reportsOut *models.ReportResponse, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	qry2 := `select count(id) from public."Reports"`
	err = DB.QueryRow(qry2).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get all reports(amount):", err)
		return nil, err
	}
	reportsOut = &models.ReportResponse{
		Reports: nil,
		Amount:  amount,
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
		reportsOut.Reports = append(reportsOut.Reports, NewReport)
	}
	return reportsOut, nil
}

func (r ReportRepository) Create(report *models.Report) (id int64, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	qry := `INSERT INTO public."Reports" (description, receiver_joke_id, sender_id, receiver_id) values ($1, $2, $3, $4) RETURNING id`
	err = DB.QueryRow(qry, report.Description, report.ReceiverJokeId, report.SenderId, report.ReceiverId).Scan(&id)
	if err != nil {
		log.Println("Error while trying to create report:", err)
		return -1, err
	}
	return id, err
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
