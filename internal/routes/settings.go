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
			Pattern:     "/profile/",
			HandlerFunc: handlers.GetUserSettingsHandler,
		},
		{
			Name:        "DevelopPage",
			Method:      "GET",
			Pattern:     "/develop/",
			HandlerFunc: handlers.GetAllReportsHandler,
		},
		{
			Name:        "DevelopPageApplyReport",
			Method:      "PUT",
			Pattern:     "/develop/apply_report/",
			HandlerFunc: handlers.ApplyReportHandler,
		},
		{
			Name:        "DevelopPageDenyReport",
			Method:      "PUT",
			Pattern:     "/develop/deny_report/",
			HandlerFunc: handlers.DenyReportHandler,
		},
	},
}
