package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var User = router.RoutePrefix{
	Prefix: "/user",
	SubRoutes: []router.Route{
		{
			Name:        "UserIndex",
			Method:      "GET",
			Pattern:     "/{id:[0-9]+}",
			HandlerFunc: handlers.UserIndexHandler,
			Protected:   false,
		},
	},
}
