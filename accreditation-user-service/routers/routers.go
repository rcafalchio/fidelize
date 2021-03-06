package routers

import (
	"fidelize/accreditation-user-service/common"
	"fidelize/accreditation-user-service/controllers"
	"log"

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
	userController := controllers.NewUserController(common.GetMongoSession())

	// Get a user by id
	r.GET("/users/:id", userController.GetUser)

	// Create a new user
	r.POST("/users", userController.CreateUser)

	log.Println("ADD ROTE /users")
	// Get all users
	r.GET("/users", userController.GetAllUsers)

}
