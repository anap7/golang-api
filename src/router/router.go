package router

import "github.com/gorilla/mux"

//It will generate a new router with all routes configured
func Generate() *mux.Router {
	return mux.NewRouter()
}