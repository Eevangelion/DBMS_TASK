package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Search = router.RoutePrefix{
	Prefix: "/search",
	SubRoutes: []router.Route{
		{
			Name:        "GetJokesWithKeyWord",
			Method:      "GET",
			Pattern:     "/keyword_search/",
			HandlerFunc: handlers.SearchHandler,
		},
		{
			Name:        "GetJokesWithTag",
			Method:      "GET",
			Pattern:     "/tag_search/",
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
