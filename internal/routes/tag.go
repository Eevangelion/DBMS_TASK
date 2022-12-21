package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Tag = router.RoutePrefix{
	Prefix: "/report",
	SubRoutes: []router.Route{
		{
			Name:        "CreateReport",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: handlers.CreateTagHandler,
		},
		{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/delete",
			HandlerFunc: handlers.DeleteTagHandler,
		},
	},
}
