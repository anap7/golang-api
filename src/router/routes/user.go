package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthenticationIsNeeded: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.GetUser,
		AuthenticationIsNeeded: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetAllUsers,
		AuthenticationIsNeeded: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthenticationIsNeeded: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthenticationIsNeeded: false,
	},
}
