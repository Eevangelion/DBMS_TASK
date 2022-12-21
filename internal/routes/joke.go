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
			Method:      "DELETE",
			Pattern:     "/delete",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
	},
}
