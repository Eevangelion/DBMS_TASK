package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Search = router.RoutePrefix{
	Prefix: "/search",
	SubRoutes: []router.Route{
		{
			Name:        "GetJokesSearch",
			Method:      "GET",
			Pattern:     "/{t}/{q}/",
			HandlerFunc: handlers.SearchHandler,
		},
	},
}
