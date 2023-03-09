package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//It will generate a new router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	//Call a config that prepare all routes
	return routes.Config(r)
}