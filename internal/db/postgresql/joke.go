package psql

import (
	"errors"
	"log"

	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/Sakagam1/DBMS_TASK/internal/repositories"
)

type JokeRepository struct {
	joke repositories.IJoke
}

func (j JokeRepository) AddToFavorite(user_id int, joke_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `INSERT INTO public."Favorite jokes" (joke_id, user_id) values ($1, $2)`
	_, err = DB.Exec(qry, user_id, joke_id)
	if err != nil {
		log.Println("Adding to favorite error:", err)
		return err
	}
	return nil
}

func (j JokeRepository) GetUserFavoriteJokes(user_id int) (jokes []models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating from public."Jokes", public."Users", public."Favorite jokes" where "Users".id="Favorite jokes".user_id and "Favorite jokes".joke_id="Jokes".id and "Users".id=$1`
	rows, err := DB.Query(qry, user_id)
	if err != nil {
		log.Println("Connection Error:", err)
		return nil, err
	}
	for rows.Next() {
		var id, rating int
		var header, description string
		err := rows.Scan(&id, &header, &description, &rating)
		if err != nil {
			log.Println("Err while scanning rows", err)
		}
		NewJoke := models.Joke{
			ID:          id,
			Header:      header,
			Description: description,
			Rating:      rating,
			AuthorId:    user_id,
		}
		jokes = append(jokes, NewJoke)
	}
	defer rows.Close()
	return jokes, nil
}

func (j JokeRepository) GetUserJokes(user_id int) (jokes []models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select "Jokes".id, "Jokes".header, "Jokes".description, "Jokes".rating from public."Jokes", public."Users" where "Users".id="Jokes".author_id and "Users".id=$1`
	rows, err := DB.Query(qry, user_id)
	if err != nil {
		log.Println("Connection Error:", err)
		return nil, err
	}
	for rows.Next() {
		var id, rating int
		var header, description string
		err := rows.Scan(&id, &header, &description, &rating)
		if err != nil {
			log.Println("Err while scanning rows", err)
		}
		NewJoke := models.Joke{
			ID:          id,
			Header:      header,
			Description: description,
			Rating:      rating,
			AuthorId:    user_id,
		}
		jokes = append(jokes, NewJoke)
	}
	defer rows.Close()
	return jokes, nil
}

func (j JokeRepository) GetJokeTags(joke_id int) (tags []models.Tag, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select "Tags".id, "Tags".name from public."Jokes", public."TagsJokes", public."Tags" where "Jokes".id="TagsJokes".joke_id and "TagsJokes".tag_id="Tags".id and "Jokes".id=$1`
	rows, err := DB.Query(qry, joke_id)
	if err != nil {
		log.Println("Connection Error:", err)
		return nil, err
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Err while scanning rows", err)
			return nil, err
		}
		NewTag := models.Tag{
			ID:   id,
			Name: name,
		}
		tags = append(tags, NewTag)
	}
	defer rows.Close()
	return tags, nil
}

func (j JokeRepository) AddTagToJoke(joke_id int, tag_id int) (err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return err
	}
	qry := `INSERT INTO public."TagsJokes" (tag_id, joke_id) values ($1, $2)`
	_, err = DB.Exec(qry, tag_id, joke_id)
	if err != nil {
		log.Println("Error while trying to add tag to joke:", err)
		return err
	}
	return nil
}

func (j JokeRepository) GetJokeByID(JokeId int) (userOut *models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Jokes" where id=$1`
	rows, err := DB.Query(qry, JokeId)
	if err != nil {
		log.Println("Searching joke by id error:", err)
	}
	var id, rating, author_id int
	var header, description string
	id = -1
	for rows.Next() {
		err := rows.Scan(&id, &header, &description, &rating, &author_id)
		if err != nil {
			log.Println("Err while scanning rows:", err)
		}
	}
	defer rows.Close()
	if id != -1 {
		return &models.Joke{
			ID:          id,
			Header:      header,
			Description: description,
			Rating:      rating,
			AuthorId:    author_id,
		}, nil
	}
	return &models.Joke{}, errors.New("Joke with this id does not exist!")
}

func (j JokeRepository) GetAll() (jokes []models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `select * from public."Jokes"`
	rows, err := DB.Query(qry)
	if err != nil {
		log.Println("Connection Error:", err)
	}
	for rows.Next() {
		var id, rating, author_id int
		var header, description string
		err := rows.Scan(&id, &header, &description, &rating, &author_id)
		if err != nil {
			log.Println("Err while scanning rows", err)
		}
		NewJoke := models.Joke{
			ID:          id,
			Header:      header,
			Description: description,
			Rating:      rating,
			AuthorId:    author_id,
		}
		jokes = append(jokes, NewJoke)
	}
	defer rows.Close()
	return jokes, nil
}

func (j JokeRepository) Create(joke *models.Joke) (jokeOut *models.Joke, err error) {
	DB, err := connection.GetConnectionToDB()
	if err != nil {
		log.Println("Connection error:", err)
		return nil, err
	}
	qry := `INSERT INTO public."Jokes" (id, header, description, rating, author_id) values ($1, $2, $3, $4, $5)`
	result, err := DB.Exec(qry, joke.ID, joke.Header, joke.Description, joke.Rating, joke.AuthorId)
	if err != nil {
		log.Println("Joke creation error:", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Joke searching while adding joke error:", err)
	}
	jokeOut, err = j.GetJokeByID(int(id))
	return jokeOut, err
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
		log.Println("Joke deletion error:", err)
		return err
	}
	return nil
}
