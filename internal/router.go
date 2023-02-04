package main

import (
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/handlers"
	customRouter "github.com/Sakagam1/DBMS_TASK/internal/router"
	"github.com/Sakagam1/DBMS_TASK/internal/routes"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	customRouter.AppRoutes = append(
		customRouter.AppRoutes,
		routes.User,
		routes.Joke,
		routes.Report,
		routes.Tag,
	)

	for _, route := range customRouter.AppRoutes {

		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc

			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)

			routePrefix.
				Path(r.Pattern).
				Handler(http.HandlerFunc(handlers.OptionsHandler)).
				Methods("OPTIONS").
				Name(r.Name)

		}
	}

	return router
}
