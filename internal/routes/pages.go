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
			Pattern:     "/?sortArg={sortArg}&pageArg={pageArg}",
			HandlerFunc: handlers.GetPageOfJokesHandler,
		},
		{
			Name:        "SendReport",
			Method:      "POST",
			Pattern:     "/?sortArg={sortArg}&pageArg={pageArg}/post_report/?joke_id={joke_id}",
			HandlerFunc: handlers.CreateReportHandler,
		},
	},
}
