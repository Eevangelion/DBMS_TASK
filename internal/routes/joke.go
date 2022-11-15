package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Joke = router.RoutePrefix{
	Prefix: "/joke",
	SubRoutes: []router.Route{
		{
			Name:        "JokeIndex",
			Method:      "GET",
			Pattern:     "/{id:[0-9]+}",
			HandlerFunc: handlers.JokeIndexHandler,
			Protected:   false,
		},
	},
}
