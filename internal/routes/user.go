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
			Name:        "GetUserSubscribedJokesByID",
			Method:      "GET",
			Pattern:     "/subscribed/",
			HandlerFunc: handlers.GetUserSubscribedJokesHandler,
		},
		{
			Name:        "GetUserFavoriteJokesByID",
			Method:      "GET",
			Pattern:     "/favorites/{id}/",
			HandlerFunc: handlers.GetUserFavoriteJokesHandler,
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
			Name:        "GetUserPageByID",
			Method:      "GET",
			Pattern:     "/{id}/",
			HandlerFunc: handlers.GetUserDataByIDHandler,
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
			Pattern:     "/oauth/{code}",
			HandlerFunc: handlers.GetGithubUser,
		},
		{
			Name:        "SubscribeToUser",
			Method:      "POST",
			Pattern:     "/subscribe/",
			HandlerFunc: handlers.SubscribeToUserHandler,
		},
		{
			Name:        "GetSubscribes",
			Method:      "GET",
			Pattern:     "/subscribed/",
			HandlerFunc: handlers.GetUserSubscribedJokesHandler,
		},
		{
			Name:        "ChangeUsername",
			Method:      "PUT",
			Pattern:     "/change_name/",
			HandlerFunc: handlers.ChangeUserNameHandler,
		},
		{
			Name:        "ChangePassword",
			Method:      "PUT",
			Pattern:     "/change_password/",
			HandlerFunc: handlers.ChangeUserPasswordHandler,
		},
		{
			Name:        "RefreshUser",
			Method:      "POST",
			Pattern:     "/refresh/",
			HandlerFunc: handlers.RefreshUser,
		},
	},
}
