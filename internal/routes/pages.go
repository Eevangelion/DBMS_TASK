package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Pages = router.RoutePrefix{
	Prefix: "/feed",
	SubRoutes: []router.Route{
		{
			Name:        "GetPages",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "SendReport",
			Method:      "POST",
			Pattern:     "/post_report/",
			HandlerFunc: handlers.CreateReportHandler,
		},
	},
}
