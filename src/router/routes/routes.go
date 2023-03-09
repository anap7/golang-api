package routes

import "net/http"

// Represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	AuthenticationIsNeeded bool
}