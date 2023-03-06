package psql

import (
	"errors"
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type TagRepository struct {
	tag repositories.ITag
}

func (t TagRepository) GetTagByID(TagID int) (tagOut *models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Tags" where id=$1`
	rows, err := DB.Query(qry, TagID)
	if err != nil {
		log.Println("Error while trying to searching tag by id:", err)
	}
	var id int
	var name string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Err while scanning rows:", err)
		}
	}
	defer rows.Close()
	if id != -1 {
		return &models.Tag{
			ID:   id,
			Name: name,
		}, nil
	}
	return &models.Tag{}, errors.New("Tag with this id does not exist!")
}

func (t TagRepository) Create(tag *models.Tag) (tagOut *models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `INSERT INTO public."Tags" (name) values ($1)`
	result, err := DB.Exec(qry, tag.ID)
	if err != nil {
		log.Println("Joke creation error:", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Joke searching while adding joke error:", err)
	}
	tagOut, err = t.GetTagByID(int(id))
	return tagOut, err
}

func (t TagRepository) Delete(tag *models.Tag) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Tags" where id=$1`
	_, err = DB.Exec(qry, tag.ID)
	if err != nil {
		log.Println("Error while trying to delete tag:", err)
		return err
	}
	return nil
}
