package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Report = router.RoutePrefix{
	Prefix: "/report",
	SubRoutes: []router.Route{
		{
			Name:        "CreateReport",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: handlers.CreateReportHandler,
		},
		{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/delete",
			HandlerFunc: handlers.DeleteReportHandler,
		},
	},
}
