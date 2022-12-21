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
			Method:      "Get",
			Pattern:     "/new/?page=<int>",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "GetTopPages",
			Method:      "Get",
			Pattern:     "/top/?t={hour | day | week | month | all}?page=<int>",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "SendReport",
			Method:      "Get",
			Pattern:     "/{new | top/?t={hour | day | week | month | all} }?page=<int>/post_report/?joke_id=<int>",
			HandlerFunc: handlers.CreateReportHandler,
		},
	},
}
