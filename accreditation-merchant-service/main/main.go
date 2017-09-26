package main

import ( // Standard packages
	// My packages
	"accreditation-merchant-service/common"
)

func main() {
	//Getting the variables of config.json
	common.InitConfigs()

	// Instantiate a new router
	/*
		r := routers.NewRouter()

		// Get the routers
		r.GetRouters()

		// Start the server
		log.Fatal(http.ListenAndServe(common.AppSettings.Server, r))
	*/
}
