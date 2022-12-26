package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Joke = router.RoutePrefix{
	Prefix: "/joke",
	SubRoutes: []router.Route{
		{
			Name:        "CreateJoke",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "Delete",
			Method:      "PUT",
			Pattern:     "/delete",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
		{
			Name:        "Get",
			Method:      "GET",
			Pattern:     "/tags?id={joke_id}",
			HandlerFunc: handlers.GetJokeTagsHandler,
		},
		{
			Name:        "AddToFavorite",
			Method:      "POST",
			Pattern:     "/addToFavorites?user_id={user_id}&joke_id={joke_id}",
			HandlerFunc: handlers.AddToFavoriteHandler,
		},
		{
			Name:        "DeleteFromFavorite",
			Method:      "POST",
			Pattern:     "/removeFromFavorites?user_id={user_id}&joke_id={joke_id}",
			HandlerFunc: handlers.DeleteFromFavoriteHandler,
		},
	},
}
