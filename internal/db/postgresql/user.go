package psql

import (
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
	if err != nil {
		log.Println("Error while searching user by id:", err)
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Err while scanning rows:", err)
		}
	}
	defer rows.Close()
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
	if err != nil {
		log.Println("Error while searching user by username:", err)
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Err while scanning rows:", err)
		}
	}
	defer rows.Close()
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
	if err != nil {
		log.Println("Error while searching user by email:", err)
	}
	var id, reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Err while scanning rows", err)
		}
	}
	defer rows.Close()
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
	unban_date := current_time.Format("02.01.2006")
	qry := `UPDATE public."Users" SET unban_date=$1 where id=$2`
	_, err = DB.Exec(qry, unban_date, user.ID)
	if err != nil {
		log.Println("User ban error:", err)
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
		log.Println("User deletion error:", err)
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
	if err != nil {
		log.Println("Connection Error:", err)
	}
	for rows.Next() {
		var id, reports, remaining_reports int
		var name, email, role, unban_date, transformed_password string
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("err while scanning rows", err)
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
	defer rows.Close()
	return users, nil
}
