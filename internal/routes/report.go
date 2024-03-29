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
			Pattern:     "/create/",
			HandlerFunc: handlers.CreateReportHandler,
		},
		{
			Name:        "DeleteReport",
			Method:      "DELETE",
			Pattern:     "/delete/",
			HandlerFunc: handlers.DeleteReportHandler,
		},
		{
			Name:        "GetReportByID",
			Method:      "GET",
			Pattern:     "/{id}/",
			HandlerFunc: handlers.GetReportByIDHandler,
		},
		{
			Name:        "ApplyReport",
			Method:      "POST",
			Pattern:     "/apply/",
			HandlerFunc: handlers.ApplyReportHandler,
		},
		{
			Name:        "GetReports",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.GetAllReportsHandler,
		},
	},
}
