package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Pages = router.RoutePrefix{
	Prefix: "/feed",
	SubRoutes: []router.Route{
		{
			Name:        "GetNewPages",
			Method:      "GET",
			Pattern:     "/new/?page=<int>",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "GetTopPages",
			Method:      "GET",
			Pattern:     "/top/?t={query}&page=<int>",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "SendReport",
			Method:      "POST",
			Pattern:     "/create_report/?joke_id={joke_id}",
			HandlerFunc: handlers.CreateReportHandler,
		},
		{
			Name:        "AddToFavorite",
			Method:      "POST",
			Pattern:     "/favorite/?joke_id={joke_id}",
			HandlerFunc: handlers.AddToFavoriteHandler,
		},
		{
			Name:        "DeleteFromFavorite",
			Method:      "POST",
			Pattern:     "/favorite/?joke_id={joke_id}",
			HandlerFunc: handlers.DeleteFromFavoriteHandler,
		},
	},
}
