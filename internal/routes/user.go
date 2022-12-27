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
			Pattern:     "/{username}",
			HandlerFunc: handlers.CreateUserHandler,
		},
		{
			Name:        "CreateJokeFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/create_joke",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "DeleteJokeFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/delete_joke/?joke_id={joke_id}",
			HandlerFunc: handlers.DeleteJokeHandler,
		},
		{
			Name:        "GetUserPage",
			Method:      "GET",
			Pattern:     "/{username}/",
			HandlerFunc: handlers.GetUserJokesHandler,
		},
		{
			Name:        "SendReportFromUserPage",
			Method:      "POST",
			Pattern:     "/{username}/create_report/?joke_id={joke_id}",
			HandlerFunc: handlers.CreateReportHandler,
		},
		{
			Name:        "GetGithubUser",
			Method:      "GET",
			Pattern:     "/oauth",
			HandlerFunc: handlers.GetGithubUser,
		},
	},
}
