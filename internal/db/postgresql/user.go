package psql

import (
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
	var amount int
	qry2 := `select count(name) from public."Users" where id=$1`
	err = DB.QueryRow(qry2, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by id:", err)
		return nil, err
	}
	if amount == 0 {
		return userOut, nil
	}
	var reports, remaining_reports int
	var name, email, role, unban_date, transformed_password string
	qry := `select name, email, reports, remaining_reports, role, unban_date, transformed_password from public."Users" where id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
	if err != nil {
		log.Println("Error while trying to get user by id:", err)
		return nil, err
	}
	return &models.User{
		ID:                  user_id,
		Name:                name,
		Email:               email,
		Reports:             reports,
		RemainingReports:    remaining_reports,
		Role:                role,
		UnbanDate:           unban_date,
		TransformedPassword: transformed_password,
	}, nil
}

func (u UserRepository) GetUserByUsername(username string) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	qry2 := `select count(id) from public."Users" where name=$1`
	err = DB.QueryRow(qry2, username).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by username:", err)
		return nil, err
	}
	if amount == 0 {
		return userOut, nil
	}
	var id, reports, remaining_reports int
	var email, role, unban_date, transformed_password string
	qry := `select id, email, reports, remaining_reports, role, unban_date, transformed_password from public."Users" where name=$1`
	err = DB.QueryRow(qry, username).Scan(&id, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
	if err != nil {
		log.Println("Error while trying to get user by username:", err)
		return nil, err
	}
	return &models.User{
		ID:                  id,
		Name:                username,
		Email:               email,
		Reports:             reports,
		RemainingReports:    remaining_reports,
		Role:                role,
		UnbanDate:           unban_date,
		TransformedPassword: transformed_password,
	}, nil
}

func (u UserRepository) GetUserByEmail(Email string) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	qry2 := `select count(id) from public."Users" where email=$1`
	err = DB.QueryRow(qry2, Email).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by email:", err)
		return nil, err
	}
	if amount == 0 {
		return userOut, nil
	}
	var id, reports, remaining_reports int
	var name, role, unban_date, transformed_password string
	qry := `select id, name, reports, remaining_reports, role, unban_date, transformed_password from public."Users" where email=$1`
	err = DB.QueryRow(qry, Email).Scan(&id, &name, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
	if err != nil {
		log.Println("Error while trying to get user by email:", err)
		return nil, err
	}
	return &models.User{
		ID:                  id,
		Name:                name,
		Email:               Email,
		Reports:             reports,
		RemainingReports:    remaining_reports,
		Role:                role,
		UnbanDate:           unban_date,
		TransformedPassword: transformed_password,
	}, nil
}

func (u UserRepository) Create(user *models.User) (id int64, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	qry := `INSERT INTO public."Users" (name, email, role, transformed_password) values ($1, $2, $3, $4) RETURNING id`
	err = DB.QueryRow(qry, user.Name, user.Email, "guest", user.TransformedPassword).Scan(&id)
	if err != nil {
		log.Println("User creation error:", err)
		return -1, err
	}
	return id, err
}

func (u UserRepository) Ban(user_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	current_time := time.Now()
	current_time = current_time.Add(time.Hour * 24 * 7)
	unban_date := current_time.Format("2006-01-02")
	qry := `UPDATE public."Users" SET unban_date=$1 where id=$2`
	_, err = DB.Exec(qry, unban_date, user_id)
	if err != nil {
		log.Println("Error while trying to ban user:", err)
		return err
	}
	return nil
}

func (u UserRepository) Delete(user_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Users" where id=$1`
	_, err = DB.Exec(qry, user_id)
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
		log.Println("Error while trying to get all users:", err)
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

func (u UserRepository) GetPeopleByKeyword(keyword string, page int, pageSize int) (users []models.UserResponseSearch, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Users" where "Users".name LIKE '%` + keyword + `%' LIMIT $1 OFFSET $2`
	rows, err := DB.Query(qry, pageSize, (page-1)*pageSize)
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
		amount, err := u.GetUserJokesCount(id)
		count, err := u.GetSubscribedPeopleCount(id)
		NewUser := models.UserResponseSearch{
			ID:               id,
			Name:             name,
			Role:             role,
			PostsCount:       amount,
			SubscribersCount: count,
		}
		users = append(users, NewUser)
	}
	return users, nil
}

func (u UserRepository) ChangeUserRemainingReports(user_sender_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `UPDATE public."Users" SET remaining_reports=remaining_reports-1 where id=$1`
	_, err = DB.Exec(qry, user_sender_id)
	if err != nil {
		log.Println("Error while trying to change user reports count:", err)
		return err
	}
	return nil
}

func (u UserRepository) ChangeUserReportsCount(user_receiver_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `UPDATE public."Users" SET reports=reports+1 where id=$1`
	_, err = DB.Exec(qry, user_receiver_id)
	if err != nil {
		log.Println("Error while trying to change user reports count:", err)
		return err
	}
	return nil
}

func (u UserRepository) ChangeUserName(user_id int, new_name string) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `UPDATE public."Users" SET name=$1 where id=$2`
	_, err = DB.Exec(qry, new_name, user_id)
	if err != nil {
		log.Println("Error while trying to change user user name:", err)
		return err
	}
	return nil
}

func (u UserRepository) ChangeUserPassword(user_id int, new_transformed_password string) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `UPDATE public."Users" SET transformed_password=$1 where id=$2`
	_, err = DB.Exec(qry, new_transformed_password, user_id)
	if err != nil {
		log.Println("Error while trying to change user reports count:", err)
		return err
	}
	return nil
}

func (u UserRepository) GetUserByGithubID(user_id int) (userOut *models.GitUser, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	qry2 := `select count(git_id) from public."GithubUsers" where git_id=$1`
	err = DB.QueryRow(qry2, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get github user by id:", err)
		return nil, err
	}
	if amount == 0 {
		return userOut, nil
	}
	qry := `select * from public."GithubUsers" where git_id=$1`
	var git_id, inner_id int
	err = DB.QueryRow(qry, user_id).Scan(&git_id, &inner_id)
	if err != nil {
		log.Println("Error while trying to get github user by id:", err)
		return nil, err
	}
	return &models.GitUser{
		Git_ID:   user_id,
		Inner_ID: inner_id,
	}, nil
}

func (u UserRepository) CreateGithubUserWithID(user_id int, inner_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `INSERT INTO public."Users" (git_id,inner_id) values ($1, $2)`
	_, err = DB.Exec(qry, user_id, inner_id)
	if err != nil {
		log.Println("Error while trying to create github user:", err)
		return err
	}
	return nil
}

func (u UserRepository) GetSubscribedPeopleCount(user_id int) (amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return 0, err
	}
	qry := `select count(receiver_id) from public."UserSubscribes" where receiver_id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&amount)
	return amount, nil
}

func (u UserRepository) GetUserJokesCount(user_id int) (amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return 0, err
	}
	qry := `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&amount)
	return amount, nil
}
