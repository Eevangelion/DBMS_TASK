package psql

import (
	"errors"
	"log"
	"time"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type UserRepository struct {
	user repositories.IUser
}

func (u UserRepository) GetUserByID(user_id int) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users" where id=$1`
	rows, err := DB.Query(qry, user_id)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get user by id:", err)
		return nil, err
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
		}
	}
	if id != -1 {
		return &models.User{
			ID:                  id,
			Name:                name,
			Email:               email,
			Reports:             reports,
			RemainingReports:    remaining_reports,
			Role:                role,
			UnbanDate:           unban_date,
			TransformedPassword: transformed_password,
		}, nil
	}
	return nil, errors.New("User with this id does not exist!")
}

func (u UserRepository) GetUserByUsername(username string) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users" where name=$1`
	rows, err := DB.Query(qry, username)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get user by username:", err)
		return nil, err
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
		}
	}
	if id != -1 {
		return &models.User{
			ID:                  id,
			Name:                name,
			Email:               email,
			Reports:             reports,
			RemainingReports:    remaining_reports,
			Role:                role,
			UnbanDate:           unban_date,
			TransformedPassword: transformed_password,
		}, nil
	}
	return nil, errors.New("User with this username does not exist!")
}

func (u UserRepository) GetUserByEmail(Email string) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users" where email=$1`
	rows, err := DB.Query(qry, Email)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get user by email:", err)
		return nil, err
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
		}
	}
	if id != -1 {
		return &models.User{
			ID:                  id,
			Name:                name,
			Email:               email,
			Reports:             reports,
			RemainingReports:    remaining_reports,
			Role:                role,
			UnbanDate:           unban_date,
			TransformedPassword: transformed_password,
		}, nil
	}
	return nil, errors.New("User with this email does not exist!")
}

func (u UserRepository) Create(user *models.User) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `INSERT INTO public."Users" (name, email, role, transformed_password) values ($1, $2, $3, $4)`
	_, err = DB.Exec(qry, user.Name, user.Email, user.Role, user.TransformedPassword)
	if err != nil {
		log.Println("User creation error:", err)
		return nil, err
	}
	userOut, err = u.GetUserByUsername(user.Name)
	return userOut, err
}

func (u UserRepository) Ban(user *models.User) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	current_time := time.Now()
	current_time = current_time.Add(time.Hour * 24 * 7)
	unban_date := current_time.Format("2006-01-02")
	qry := `UPDATE public."Users" SET unban_date=$1 where id=$2`
	_, err = DB.Exec(qry, unban_date, user.ID)
	if err != nil {
		log.Println("Error while trying to ban user:", err)
		return err
	}
	return nil
}

func (u UserRepository) Delete(user *models.User) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Users" where name=$1`
	_, err = DB.Exec(qry, user.Name)
	if err != nil {
		log.Println("Error while trying to delete user:", err)
		return err
	}
	return nil
}

func (u UserRepository) GetAll() (users []models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users"`
	rows, err := DB.Query(qry)
	defer rows.Close()
	if err != nil {
		log.Println("Connection Error:", err)
		return nil, err
	}
	for rows.Next() {
		var id, reports, remaining_reports int
		var name, email, role, unban_date, transformed_password string
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
		NewUser := models.User{
			ID:                  id,
			Name:                name,
			Email:               email,
			Reports:             reports,
			RemainingReports:    remaining_reports,
			Role:                role,
			UnbanDate:           unban_date,
			TransformedPassword: transformed_password,
		}
		users = append(users, NewUser)
	}
	return users, nil
}

func (u UserRepository) GetPeopleByKeyWord(keyword string) (users []models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users" where "Users".name LIKE '$1'`
	rows, err := DB.Query(qry, keyword)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get people by keyword:", err)
		return nil, err
	}
	for rows.Next() {
		var id, reports, remaining_reports int
		var name, email, role, unban_date, transformed_password string
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
		NewUser := models.User{
			ID:                  id,
			Name:                name,
			Email:               email,
			Reports:             reports,
			RemainingReports:    remaining_reports,
			Role:                role,
			UnbanDate:           unban_date,
			TransformedPassword: transformed_password,
		}
		users = append(users, NewUser)
	}
	return users, nil
}
