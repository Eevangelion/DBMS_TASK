package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var User = router.RoutePrefix{
	Prefix: "/user",
	SubRoutes: []router.Route{
		{
			Name:        "CreateUser",
			Method:      "POST",
			Pattern:     "/create_user/",
			HandlerFunc: handlers.CreateUserHandler,
		},
		{
			Name:        "GetUserJokes",
			Method:      "GET",
			Pattern:     "/{username}/jokes/",
			HandlerFunc: handlers.GetUserJokesHandler,
		},
		{
			Name:        "GetUserPageByName",
			Method:      "GET",
			Pattern:     "/{username}/data/",
			HandlerFunc: handlers.GetUserDataHandler,
		},
		{
			Name:        "GetGithubUser",
			Method:      "GET",
			Pattern:     "/oauth/",
			HandlerFunc: handlers.GetGithubUser,
		},
	},
}
