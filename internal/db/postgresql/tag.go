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
	var amount int
	var name string
	qry := `select name from public."Tags" where id=$1`
	qry_count := `select count(name) from public."Tags" where id=$1`
	err = DB.QueryRow(qry_count, tag_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get tag by ID (amount):", err)
		return tagOut, err
	}
	if amount == 0 {
		return tagOut, nil
	}
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
	var amount int
	qry_count := `select count(id) from public."Tags" where name=$1`
	err = DB.QueryRow(qry_count, tag_name).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to create tag (amount):", err)
		return -1, err
	}
	if amount != 0 {
		log.Println("Error while trying to create tag: tag already exist")
		return -1, nil
	}
	qry := `INSERT INTO public."Tags" (name) values ($1) RETURNING id`
	err = DB.QueryRow(qry, tag_name).Scan(&id)
	if err != nil {
		log.Println("Error while trying to create tag:", err)
		return -1, err
	}
	return id, err
}

func (t TagRepository) Delete(tag_name string) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Tags" where name=$1`
	_, err = DB.Exec(qry, tag_name)
	if err != nil {
		log.Println("Error while trying to delete tag:", err)
		return err
	}
	return nil
}

func (t TagRepository) GetAllTags() (tagsOut []models.Tag, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry2 := `select count(id) from public."Tags"`
	err = DB.QueryRow(qry2).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get all tags(amount):", err)
		return nil, -1, err
	}
	qry := `select * from public."Tags"`
	rows, err := DB.Query(qry)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get all tags:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
		}
		NewTag := models.Tag{
			ID:   id,
			Name: name,
		}
		tagsOut = append(tagsOut, NewTag)
	}
	return tagsOut, amount, nil
}
