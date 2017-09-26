package routers

import (
	"accreditation-merchant-service/common"
	"accreditation-merchant-service/controllers"
	//"github.com/julienschmidt/httprouter"
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
	uc := controllers.NewUserController(common.GetMongoSession())
	wc := controllers.NewWidgetController(common.GetMongoSession())
	ac := controllers.NewAuthController()

	// Get a user by id
	r.GET("/users/:id", uc.GetUser)

	// Create a new user
	r.POST("/users", uc.CreateUser)

	// Get all users
	r.GET("/users", uc.GetAllUsers)

	// Get a widget by id
	r.GET("/widgets/:id", wc.GetWidget)

	// Create a new widget
	r.POST("/widgets", wc.CreateWidget)

	// Get all widgets
	r.GET("/widgets", wc.GetAllWidgets)

	// Remove a existing widget
	r.DELETE("/widgets/:id", wc.RemoveWidget)

	// Update a existing widget
	r.PUT("/widgets/:id", wc.UpdateWidget)

	r.GET("/authenticate", ac.GetToken)

}
