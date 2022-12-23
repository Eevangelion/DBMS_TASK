package psql

import (
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type TagRepository struct {
	tag repositories.ITag
}

func (t TagRepository) GetTagByID(tag_id int) (tagOut *models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var name string
	qry := `select name from public."Tags" where id=$1`
	err = DB.QueryRow(qry, tag_id).Scan(&name)
	if err != nil {
		log.Println("Error while trying to get tag by id:", err)
	}
	return &models.Tag{
		ID:   tag_id,
		Name: name,
	}, nil

}

func (t TagRepository) Create(tag_name string) (id int64, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	qry := `INSERT INTO public."Tags" (name) values ($1) RETURNING id`
	err = DB.QueryRow(qry, tag_name).Scan(&id)
	if err != nil {
		log.Println("Error while trying to create tag:", err)
		return -1, err
	}
	return id, err
}

func (t TagRepository) Delete(tag_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Tags" where id=$1`
	_, err = DB.Exec(qry, tag_id)
	if err != nil {
		log.Println("Error while trying to delete tag:", err)
		return err
	}
	return nil
}

func (t TagRepository) GetAll() (tags []models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Tags"`
	rows, err := DB.Query(qry)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get all tags:", err)
		return nil, err
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
		NewTag := models.Tag{
			ID:   id,
			Name: name,
		}
		tags = append(tags, NewTag)
	}
	return tags, nil
}
