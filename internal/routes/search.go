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
			Method:      "Get",
			Pattern:     "/?q=<string>?t=keyword",
			HandlerFunc: handlers.SearchJokesByKeywordHandler,
		},
		{
			Name:        "Get",
			Method:      "GetJokesWithTag",
			Pattern:     "/?q=<string>?t=tag",
			HandlerFunc: handlers.SearchJokesByTagHandler,
		},
		{
			Name:        "SendReport",
			Method:      "Get",
			Pattern:     "/?q=<string>?t=people",
			HandlerFunc: handlers.SearchPeopleHandler,
		},
	},
}
