package psql

import (
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type JokeRepository struct {
	joke repositories.IJoke
}

func (j JokeRepository) SubscribeToUser(receiver_id int, sender_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	var amount int
	qry_count := `select count(receiver_id) from public."UserSubscribes" where receiver_id=$1 and sender_id=$2`
	err = DB.QueryRow(qry_count, receiver_id, sender_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to subscribe to user (amount):", err)
		return err
	}
	if amount != 0 {
		log.Println("Error while trying to subscribe to user: relation already exist")
		return nil
	}
	qry := `INSERT INTO public."UserSubscribes" (receiver_id, sender_id) values ($1, $2)`
	_, err = DB.Exec(qry, receiver_id, sender_id)
	if err != nil {
		log.Println("Error while trying to subscribe to user:", err)
		return err
	}
	return nil
}

func (j JokeRepository) UnSubscribeFromUser(receiver_id int, sender_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."UserSubscribes" where receiver_id=$1 and sender_id=$2`
	_, err = DB.Exec(qry, receiver_id, sender_id)
	if err != nil {
		log.Println("Error while trying to UnSubscribe from user:", err)
		return err
	}
	return nil
}

func (j JokeRepository) GetUserSubscribedJokes(user_id int, page int, pageSize int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 ORDER BY creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1`
	}
	if sort_mode == "alltime" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1`
	}
	if sort_mode == "hour" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `SELECT "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `SELECT count("Jokes".id) FROM public."Users", public."UserSubscribes", public."Jokes" where "Users".id="UserSubscribes".receiver_id and "Users".id="Jokes".author_id and "UserSubscribes".sender_id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user subscribed jokes (amount):", err)
		return nil, -1, err
	}
	rows, err := DB.Query(qry, user_id, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get user subscribed jokes:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, rating, author_id int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &creation_date, &author_id)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     author_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) AddToFavorite(user_id int, joke_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	var amount int
	qry_count := `select count(user_id) from public."Favorite jokes" where user_id=$1 and joke_id=$2`
	err = DB.QueryRow(qry_count, user_id, joke_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to add to favorite (amount):", err)
		return err
	}
	if amount != 0 {
		log.Println("Error while trying to add to favorite: relation already exist")
		return nil
	}
	qry := `INSERT INTO public."Favorite jokes" (user_id, joke_id) values ($1, $2)`
	qry2 := `UPDATE public."Jokes" SET rating=rating + 1 WHERE id=$1`
	_, err = DB.Exec(qry, user_id, joke_id)
	if err != nil {
		log.Println("Error while trying to add to favorite:", err)
		return err
	}
	_, err = DB.Exec(qry2, joke_id)
	if err != nil {
		log.Println("Error while trying to add to favorite:", err)
		return err
	}
	return nil
}

func (j JokeRepository) DeleteFromFavorite(user_id int, joke_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Favorite jokes" where user_id=$1 and joke_id=$2`
	qry_count := `UPDATE public."Jokes" SET rating=rating - 1 WHERE id=$1`
	_, err = DB.Exec(qry, user_id, joke_id)
	if err != nil {
		log.Println("Error while trying to delete from favorite:", err)
		return err
	}
	_, err = DB.Exec(qry_count, joke_id)
	if err != nil {
		log.Println("Error while trying to delete from favorite:", err)
		return err
	}
	return nil
}

func (j JokeRepository) CheckIfInFavorite(user_id int, joke_id int) (amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	qry := `select count(user_id) FROM public."Favorite jokes" where user_id=$1 and joke_id=$2`
	err = DB.QueryRow(qry, user_id, joke_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to delete from favorite:", err)
		return -1, err
	}
	return amount, nil
}

func (j JokeRepository) GetUserFavoriteJokes(user_id int, page int, pageSize int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 ORDER BY creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1`
	}
	if sort_mode == "alltime" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1`
	}
	if sort_mode == "hour" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date, "Jokes".author_id from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get user favorite jokes (amount):", err)
		return nil, -1, err
	}
	rows, err := DB.Query(qry, user_id, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get user favorite jokes:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, rating, author_id int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &creation_date, &author_id)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, 1, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     author_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) GetJokesByTag(tag_name string, page int, pageSize int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, 0, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 ORDER BY creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1`
	}
	if sort_mode == "alltime" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1`
	}
	if sort_mode == "hour" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Tags".name=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count, tag_name).Scan(&amount)
	if err != nil {
		log.Println("Error while getting jokes by tag (amount):", err)
		return nil, 0, err
	}
	rows, err := DB.Query(qry, tag_name, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while getting jokes by tag:", err)
		return nil, 0, err
	}
	for rows.Next() {
		var id, rating, user_id int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &user_id, &creation_date)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, 0, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     user_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) GetJokesByKeyword(keyword string, page int, pageSize int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, 0, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' ORDER BY creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count(id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%'`
	}
	if sort_mode == "alltime" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%'`
	}
	if sort_mode == "hour" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `select * from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where lower(header) LIKE '%` + keyword + `%' or lower(description) LIKE '%` + keyword + `%' and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count).Scan(&amount)
	if err != nil {
		log.Println("Error while getting jokes by keyword (amount):", err)
		return nil, 0, err
	}
	rows, err := DB.Query(qry, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while getting jokes by keyword:", err)
		return nil, 0, err
	}
	for rows.Next() {
		var id, rating, user_id int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &user_id, &creation_date)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, 0, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     user_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) GetUserJokes(user_id int, page int, pageSize int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 ORDER BY creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1`
	}
	if sort_mode == "alltime" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1`
	}
	if sort_mode == "hour" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(HOUR from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".creation_date from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $2 OFFSET $3`
		qry_count = `select count("Jokes".id) from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - "Jokes".creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count, user_id).Scan(&amount)
	if err != nil {
		log.Println("Error while getting user jokes (amount):", err)
		return nil, -1, err
	}
	rows, err := DB.Query(qry, user_id, pageSize, (page-1)*pageSize)
	defer rows.Close()
	if err != nil {
		log.Println("Error while getting user jokes:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, rating int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &creation_date)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     user_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) GetJokeTags(joke_id int) (tags []models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select "Tags".id, "Tags".name from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Jokes".id=$1`
	rows, err := DB.Query(qry, joke_id)
	defer rows.Close()
	if err != nil {
		log.Println("Error while getting joke tags:", err)
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

func (j JokeRepository) AddTagToJoke(joke_id int, tag_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	var amount int
	qry_count := `select count(tag_id) from public."TagsJokes" where joke_id=$1 and tag_id=$2`
	err = DB.QueryRow(qry_count, joke_id, tag_id).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to add tag to joke (amount):", err)
		return err
	}
	if amount != 0 {
		log.Println("Error while trying to add tag to joke: relation already exist")
		return nil
	}
	qry := `INSERT INTO public."TagsJokes" (tag_id, joke_id) values ($1, $2)`
	_, err = DB.Exec(qry, tag_id, joke_id)
	if err != nil {
		log.Println("Error while trying to add tag to joke:", err)
		return err
	}
	return nil
}

func (j JokeRepository) DeleteTagFromJoke(joke_id int, tag_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."TagsJokes" where tag_id=$1 and joke_id=$2`
	_, err = DB.Exec(qry, tag_id, joke_id)
	if err != nil {
		log.Println("Error while trying to delete tag from joke:", err)
		return err
	}
	return nil
}

func (j JokeRepository) GetJokeByID(joke_id int) (jokeOut *models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}

	var amount int
	var rating, author_id int
	var header, description, creation_date string
	qry := `select "Jokes".header, "Jokes".description, "Jokes".rating, "Jokes".author_id, "Jokes".creation_date from public."Jokes" where id=$1`
	qry_count := `select count("Jokes".header) from public."Jokes" where id=$1`
	err = DB.QueryRow(qry_count, joke_id).Scan(&amount)
	if err != nil {
		log.Println("Error while searching joke by id (amount):", err)
	}
	if amount == 0 {
		return jokeOut, nil
	}
	err = DB.QueryRow(qry, joke_id).Scan(&header, &description, &rating, &author_id, &creation_date)
	if err != nil {
		log.Println("Error while searching joke by id:", err)
	}
	NewJoke := models.Joke{
		ID:           joke_id,
		Header:       header,
		Description:  description,
		Rating:       rating,
		AuthorId:     author_id,
		CreationDate: creation_date,
	}
	return &NewJoke, nil
}

func (j JokeRepository) GetPageOfJokes(page int, per_page int, sort_mode string) (jokes []models.Joke, amount int, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, -1, err
	}
	qry := ``
	qry_count := ``
	if sort_mode == "new" {
		qry = `select * from public."Jokes" ORDER BY creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select COUNT ("Jokes".id) from public."Jokes"`
	}
	if sort_mode == "alltime" {
		qry = `select * from public."Jokes" ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select COUNT ("Jokes".id) from public."Jokes"`
	}
	if sort_mode == "hour" {
		qry = `select * from public."Jokes" where EXTRACT(HOUR from (CURRENT_TIMESTAMP - creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) < 1 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where EXTRACT(HOUR from (CURRENT_TIMESTAMP - creation_date)) <= 1 and EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) < 1`
	}
	if sort_mode == "day" {
		qry = `select * from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 1 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 1`
	}
	if sort_mode == "week" {
		qry = `select * from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 7 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 7`
	}
	if sort_mode == "month" {
		qry = `select * from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 31 ORDER BY rating DESC, creation_date DESC LIMIT $1 OFFSET $2`
		qry_count = `select count("Jokes".id) from public."Jokes" where EXTRACT(DAY from (CURRENT_TIMESTAMP - creation_date)) <= 31`
	}
	err = DB.QueryRow(qry_count).Scan(&amount)
	if err != nil {
		log.Println("Error while trying to get page of jokes (amount):", err)
		return nil, -1, err
	}
	rows, err := DB.Query(qry, per_page, per_page*(page-1))
	defer rows.Close()
	if err != nil {
		log.Println("Error while trying to get page of jokes:", err)
		return nil, -1, err
	}
	for rows.Next() {
		var id, rating, author_id int
		var header, description, creation_date string
		err := rows.Scan(&id, &header, &description, &rating, &author_id, &creation_date)
		if err != nil {
			log.Println("Error while scanning rows:", err)
			return nil, -1, err
		}
		NewJoke := models.Joke{
			ID:           id,
			Header:       header,
			Description:  description,
			Rating:       rating,
			AuthorId:     author_id,
			CreationDate: creation_date,
		}
		jokes = append(jokes, NewJoke)
	}
	return jokes, amount, nil
}

func (j JokeRepository) Create(joke *models.Joke) (id int64, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return -1, err
	}
	qry := `INSERT INTO public."Jokes" (header, description, author_id) values ($1, $2, $3) RETURNING id`
	err = DB.QueryRow(qry, joke.Header, joke.Description, joke.AuthorId).Scan(&id)
	if err != nil {
		log.Println("Error while trying to create joke:", err)
		return -1, err
	}
	return id, nil
}

func (j JokeRepository) Delete(joke_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `DELETE FROM public."Jokes" where id=$1`
	_, err = DB.Exec(qry, joke_id)
	if err != nil {
		log.Println("Error while trying to delete joke:", err)
		return err
	}
	return nil
}
