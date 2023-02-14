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
			Pattern:     "/{type}/{query}/",
			HandlerFunc: handlers.SearchHandler,
		},
		{
			Name:        "SendReport",
			Method:      "POST",
			Pattern:     "/send_report/",
			HandlerFunc: handlers.SearchHandler,
		},
	},
}
