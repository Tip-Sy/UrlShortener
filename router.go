package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name	string
	Path	string
	Method	string
	Handler	http.HandlerFunc
}

func MyRouter(shortener *UrlShortener) *mux.Router {
	
	// Define the routes
	var routes = [...]Route{
		Route {
			"CreateShortlink",
			"/shortlink/",
			"POST",
			shortener.CreateShortlink,
		},
		Route {
			"Redirection",
			"/{shortlink}",
			"GET",
			shortener.Redirect,
		},
		Route {
			"Monitoring",
			"/admin/",
			"POST",
			shortener.Monitor,
		},
	}
	
	// Creation of a mux router
	router := mux.NewRouter().StrictSlash(true)
	
	// Add each route to the router
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}
	
	return router
}
