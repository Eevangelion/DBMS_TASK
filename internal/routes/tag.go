package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Tag = router.RoutePrefix{
	Prefix: "/tag",
	SubRoutes: []router.Route{
		{
			Name:        "TagIndex",
			Method:      "GET",
			Pattern:     "/{id:[0-9]+}",
			HandlerFunc: handlers.TagIndexHandler,
			Protected:   false,
		},
	},
}
