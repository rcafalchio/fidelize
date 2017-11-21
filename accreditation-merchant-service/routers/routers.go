package routers

import (
	"fidelize/accreditation-merchant-service/common"
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

	// Get a Controller instance
	merchantController := controllers.NewMerchantController(common.GetMongoSession())
	//ac := controllers.NewAuthController()
	// Get a user by id
	r.GET("/merchants/:id", merchantController.GetMerchant)
	// Create a new user
	r.POST("/merchants", merchantController.CreateMerchant)
	// Get all users
	r.GET("/merchants", merchantController.GetAllMerchants)
	// Disable user users
	r.POST("/merchants/disable/:id", merchantController.InactivateMerchant)

}
