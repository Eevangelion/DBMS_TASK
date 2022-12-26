package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Settings = router.RoutePrefix{
	Prefix: "/settings",
	SubRoutes: []router.Route{
		{
			Name:        "Profile",
			Method:      "GET",
			Pattern:     "/profile",
			HandlerFunc: handlers.GetUserSettingsHandler,
		},
		{
			Name:        "DevelopPage",
			Method:      "GET",
			Pattern:     "/develop",
			HandlerFunc: handlers.GetAllReportsHandler,
		},
		{
			Name:        "DevelopPageApplyReport",
			Method:      "GET",
			Pattern:     "/develop/apply_report?report_id={report_id}",
			HandlerFunc: handlers.ApplyReportHandler,
		},
		{
			Name:        "DevelopPageDenyReport",
			Method:      "GET",
			Pattern:     "/develop/deny_report?report_id={report_id}",
			HandlerFunc: handlers.DenyReportHandler,
		},
	},
}
