package routers

import (
	"fidelize/campaign-service/common"
	"fidelize/campaign-service/controllers"
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
	campaignController := controllers.NewCampaignController(common.GetMongoSession())
	ac := controllers.NewAuthController()

	// Get a user by id
	//r.GET("/users/:id", userController.GetUser)

	// Create a new user
	r.POST("/campaign", campaignController.CreateCampaign)

	log.Println("ADD ROTE /campaigns")
	// Get all users
	r.GET("/campaigns", userController.GetAllUsers)

}
