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
			Pattern:     "/create/",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/delete/",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
		{
			Name:        "GetJokeByID",
			Method:      "GET",
			Pattern:     "/{id}/",
			HandlerFunc: handlers.GetJokeByIDHandler,
		},
		{
			Name:        "Get",
			Method:      "GET",
			Pattern:     "/tags/{id}/",
			HandlerFunc: handlers.GetJokeTagsHandler,
		},
		{
			Name:        "AddToFavorite",
			Method:      "POST",
			Pattern:     "/addToFavorites/",
			HandlerFunc: handlers.AddToFavoriteHandler,
		},
		{
			Name:        "DeleteFromFavorite",
			Method:      "DELETE",
			Pattern:     "/removeFromFavorites/",
			HandlerFunc: handlers.DeleteFromFavoriteHandler,
		},
		{
			Name:        "AddTagToJoke",
			Method:      "POST",
			Pattern:     "/addTag/",
			HandlerFunc: handlers.AddTagToJokeHandler,
		},
		{
			Name:        "RemoveTagFromJoke",
			Method:      "POST",
			Pattern:     "/removeTag/",
			HandlerFunc: handlers.DeleteTagFromJokeHandler,
		},
	},
}

