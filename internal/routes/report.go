package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Report = router.RoutePrefix{
	Prefix: "/report",
	SubRoutes: []router.Route{
		{
			Name:        "ReportIndex",
			Method:      "GET",
			Pattern:     "/{id:[0-9]+}",
			HandlerFunc: handlers.ReportIndexHandler,
			Protected:   false,
		},
	},
}
