package repositories

import (
	"github.com/Sakagam1/DBMS_TASK/internal/models"
)

type IJoke interface {
	SubscribeToUser(receiver_id int, sender_id int) (err error)
	UnSubscribeFromUser(receiver_id int, sender_id int) (err error)
	GetUserSubribedJokes(user_id int, page int, pageSize int, sort_mode string) (jokes []models.Joke, err error)
	AddToFavorite(user_id int, joke_id int) (err error)
	DeleteFromFavorite(user_id int, joke_id int) (err error)
	AddTagToJoke(joke_id int, tag_id int) (err error)
	DeleteTagFromJoke(joke_id int, tag_id int) (err error)
	GetUserFavoriteJokes(user_id int, page int, pageSize int, sort_mode string) (jokes []models.Joke, err error)
	GetJokeTags(joke_id int) (tags []models.Tag, err error)
	GetUserJokes(user_id int, page int, pageSize int, sort_mode string) (jokeOut []models.Joke, err error)
	GetJokeByID(joke_id int) (userOut *models.Joke, err error)
	GetPageOfJokes(page int, pageSize int, sort_mode string) (jokes []models.Joke, err error)
	GetJokesByTag(tag_name string, page int, pageSize int, sort_mode string) (jokes []models.Joke, err error)
	GetJokesByKeyword(keyword string, page int, pageSize int, sort_mode string) (jokes []models.Joke, err error)
	Create(joke *models.Joke) (id int64, err error)
	Delete(joke_id int) (err error)
}
