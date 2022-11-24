package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var User = router.RoutePrefix{
	Prefix: "/user",
	SubRoutes: []router.Route{
		{
			Name:        "UserName",
			Method:      "GET",
			Pattern:     "/{username}",
			HandlerFunc: handlers.UserNameHandler,
		},
	},
}
