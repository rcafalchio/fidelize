package routers

import (
	"fidelize/accreditation-merchant-service/controllers"

	"github.com/julienschmidt/httprouter"
)

// Router struct represents the route
type Router struct {
	*httprouter.Router
}

// NewRouter Gets a reference to Router with a new referenced httprouter
func NewRouter() *Router {
	return &Router{httprouter.New()}
}

//GetRouters retrive all routers configured
func (r Router) GetRouters() {

	// Get a UserController instance
	tokenController := controllers.NewAuthController()

	// Get a user by id
	//r.GET("/users/:id", userController.GetUser)

	// Create a token by user
	r.POST("/token/create", tokenController.Create)
	// Create a new user
	r.GET("/authenticate", ac.GetToken)

}
