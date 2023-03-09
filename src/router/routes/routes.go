package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	AuthenticationIsNeeded bool
}

//Insert all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}