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
	if err != nil {
		log.Println("Searching joke by id error:", err)
	}
	var id, receiver_joke_id, sender_id, receiver_id int
	var description string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &description, &receiver_joke_id, &sender_id, &receiver_id)
		if err != nil {
			log.Println("Err while scanning rows:", err)
		}
	}
	defer rows.Close()
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

func (r ReportRepository) Create(report *models.Report) (reportOut *models.Report, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `INSERT INTO public."Reports" (header, description, rating, author_id) values ($1, $2, $3, $4)`
	result, err := DB.Exec(qry, report.Description, report.ReceiverJokeId, report.SenderId, report.ReceiverId)
	if err != nil {
		log.Println("Report creation error:", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Joke searching while adding joke error:", err)
	}
	reportOut, err = r.GetReportByID(int(id))
	return reportOut, err
}

func (r ReportRepository) Delete(report *models.Report) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Reports" where id=$1`
	_, err = DB.Exec(qry, report.ID)
	if err != nil {
		log.Println("Error while trying to delete report:", err)
		return err
	}
	return nil
}
