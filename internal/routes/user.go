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
			Name:        "CreateJokeFromUser",
			Method:      "POST",
			Pattern:     "/{username}/create_joke",
			HandlerFunc: handlers.CreateJokeHandler,
		},
		{
			Name:        "GetUserPage",
			Method:      "GET",
			Pattern:     "/{username}/?sort=new",
			HandlerFunc: handlers.GetUserJokesHandler,
		},
		{
			Name:        "GetUserPageSort",
			Method:      "GET",
			Pattern:     "/{username}/?sort=top?t={hour | day | week | month | all}",
			HandlerFunc: handlers.GetUserJokesHandler,
		},
		{
			Name:        "SendReportFromUserPage",
			Method:      "GET",
			Pattern:     "/{username}/{new/ | top/?t={hour | day | week | month | all} }?page=<int>/post_report/?joke_id=<int>",
			HandlerFunc: handlers.CreateReportHandler,
		},
	},
}
