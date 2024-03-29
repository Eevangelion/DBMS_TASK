package psql

import (
	"log"
	"strings"
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
	qry_count := `select count(name) from public."Users" where id=$1`
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by id (amount):", err)
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

func (u UserRepository) GetUserUnbanDate(user_id int) (unban_date string, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return "", err
	}
	var amount int
	qry_count := `select count(name) from public."Users" where id=$1`
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by id (amount):", err)
		return "", err
	}
	if amount == 0 {
		return "", nil
	}
	qry := `select unban_date from public."Users" where id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&unban_date)
	if err != nil {
		log.Println("Error while trying to get user by id:", err)
		return "", err
	}
	return unban_date, err
}

func (u UserRepository) GetUserByUsername(username string) (userOut *models.User, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	var amount int
	qry_count := `select count(id) from public."Users" where name=$1`
	err = DB.QueryRow(qry_count, username).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by username (amount):", err)
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
	qry_count := `select count(id) from public."Users" where email=$1`
	err = DB.QueryRow(qry_count, Email).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user by email (amount):", err)
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

func (u UserRepository) CreateUser(user *models.UserRequestRegister) (id int64, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	var amount int
	qry_count := `select count(id) from public."Users" where name=$1`
	err = DB.QueryRow(qry_count, user.Name).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to create user (amount):", err)
		return -1, err
	}
	if amount != 0 {
		log.Println("Error while trying to create user: user already exist")
		return -1, nil
	}
	qry := `INSERT INTO public."Users" (name, email, role, transformed_password) values ($1, $2, $3, $4) RETURNING id`
	err = DB.QueryRow(qry, user.Name, user.Email, "guest", user.Password).Scan(&id)
	if err != nil {
		log.Println("Error while trying to create user:", err)
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

func (u UserRepository) GetAll() (users []models.User, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry_count := `select count(id) from public."Users"`
	err = DB.QueryRow(qry_count).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get all users (amount):", err)
		return nil, -1, err
	}
	if amount == 0 {
		return nil, 0, nil
	}
	qry := `select * from public."Users"`
	rows, err := DB.Query(qry)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get all users:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, reports, remaining_reports int
		var name, email, role, unban_date, transformed_password string
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
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
	return users, amount, nil
}

func (u UserRepository) GetPeopleByKeyword(keyword string, page int, pageSize int) (users []models.UserResponseSearch, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry_count := `select count(id) from public."Users" where lower("Users".name) LIKE '%` + strings.ToLower(keyword) + `%'`
	err = DB.QueryRow(qry_count).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get people by keyword (amount):", err)
		return nil, -1, err
	}
	if amount == 0 {
		return nil, 0, nil
	}
	qry := `select * from public."Users" where lower("Users".name) LIKE '%` + strings.ToLower(keyword) + `%' LIMIT $1 OFFSET $2`
	rows, err := DB.Query(qry, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get people by keyword:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, reports, remaining_reports int
		var name, email, role, unban_date, transformed_password string
		err := rows.Scan(&id, &name, &email, &reports, &remaining_reports, &role, &unban_date, &transformed_password)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
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
	return users, amount, nil
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
		log.Println("Error while trying to change remaining user reports:", err)
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
		log.Println("Error while trying to change user name:", err)
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
		log.Println("Error while trying to change user password:", err)
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
	qry_count := `select count(git_id) from public."GithubUsers" where git_id=$1`
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get github user by id (amount):", err)
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
	qry := `INSERT INTO public."GithubUsers" (git_id,inner_id) values ($1, $2)`
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
	if err != nil {
		log.Println("Error while trying to get subscribed people count:", err)
		return 0, err
	}
	return amount, nil
}

func (u UserRepository) GetWhomUserSubscribedToCount(user_id int) (amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return 0, err
	}
	qry := `select count(receiver_id) from public."UserSubscribes" where sender_id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get subscribed people count:", err)
		return 0, err
	}
	return amount, nil
}

func (u UserRepository) GetWhomUserSubscribedTo(user_id int) (users []int, err error) {

	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select receiver_id from public."UserSubscribes" where sender_id=$1`
	rows, err := DB.Query(qry, user_id)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get subscribed people:", err)
		return nil, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, err
		}
		users = append(users, id)
	}
	return users, nil
}

func (u UserRepository) GetCheckIfUserSubscribed(sender_id int, receiver_id int) (check bool, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return false, err
	}
	var amount int
	qry := `select count(*) from public."UserSubscribes" where sender_id=$1 and receiver_id=$2`
	err = DB.QueryRow(qry, sender_id, receiver_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get check if user subscribed to another:", err)
		return false, err
	}
	return amount != 0, nil
}

func (u UserRepository) GetUserJokesCount(user_id int) (amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return 0, err
	}
	qry := `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1`
	err = DB.QueryRow(qry, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user jokes count:", err)
		return 0, err
	}
	return amount, nil

}

func (u UserRepository) SetRemainingReports() {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return
	}
	qry := `UPDATE public."Users" SET remaining_reports=3`
	_, err = DB.Exec(qry)
	if err != nil {
		log.Println("Error while trying to change remaining users reports:", err)
		return
	}
	return
}
