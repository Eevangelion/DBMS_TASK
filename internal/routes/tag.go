package routes

import (
	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	"github.com/Sakagam1/DBMS_TASK/internal/router"
)

var Tag = router.RoutePrefix{
	Prefix: "/tag",
	SubRoutes: []router.Route{
		{
			Name:        "CreateTag",
			Method:      "POST",
			Pattern:     "/create/",
			HandlerFunc: handlers.CreateTagHandler,
		},
		{
			Name:        "DeleteTag",
			Method:      "DELETE",
			Pattern:     "/delete/",
			HandlerFunc: handlers.DeleteTagHandler,
		},
		{
			Name:        "GetAllTags",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.GetAllTagsHandler,
		},
	},
}
