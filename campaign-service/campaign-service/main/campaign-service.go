package main

import ( // Standard packages
	// My packages

	"log"
	"net/http"

	"fidelize/accreditation-merchant-service/common"
	"fidelize/accreditation-merchant-service/routers"
)

func main() {
	log.Println("Starting accreditation merchant service...")
	//Getting the variables of config.json
	common.InitConfigs()
	// Instantiate a new router
	r := routers.NewRouter()
	// Get the routers
	r.GetRouters()
	log.Println("http.ListenAndServe at", common.AppSettings.Server)
	// Start the server
	log.Fatal(http.ListenAndServe(common.AppSettings.Server, r))
	log.Println("Stopping accreditation merchant service...")
}
