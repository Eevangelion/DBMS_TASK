package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Joke = router.RoutePrefix{
	Prefix: "/joke",
	SubRoutes: []router.Route{
		{
			Name:        "Create",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/delete",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
		{
			Name:        "GetAll",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.GetJokes,
		},
	},
}
