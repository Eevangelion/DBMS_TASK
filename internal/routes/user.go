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
			Name:        "CreateJokeFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/create_joke/",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "DeleteJokeFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/delete_joke/",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
		{
			Name:        "GetUserJokes",
			Method:      "GET",
			Pattern:     "/{username}/jokes/",
			HandlerFunc: handlers.GetUserJokesHandler,
		},
		{
			Name:        "GetUserPage",
			Method:      "GET",
			Pattern:     "/{username}/data/",
			HandlerFunc: handlers.GetUserDataHandler,
		},
		{
			Name:        "SendReportFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/create_report/",
			HandlerFunc: handlers.CreateReportHandler,
		},
		{
			Name:        "GetGithubUser",
			Method:      "GET",
			Pattern:     "/oauth/",
			HandlerFunc: handlers.GetGithubUser,
		},
	},
}
