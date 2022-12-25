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
			Pattern:     "/?q={query}?t=keyword",
			HandlerFunc: handlers.SearchJokesByKeywordHandler,
		},
		{
			Name:        "GetJokesWithTag",
			Method:      "GET",
			Pattern:     "/?q={query}}?t=tag",
			HandlerFunc: handlers.SearchJokesByTagHandler,
		},
		{
			Name:        "SendReport",
			Method:      "Get",
			Pattern:     "/?q={query}}?t=people",
			HandlerFunc: handlers.SearchPeopleHandler,
		},
	},
}
