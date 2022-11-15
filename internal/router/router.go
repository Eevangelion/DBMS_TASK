package router

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

var AppRoutes []RoutePrefix
