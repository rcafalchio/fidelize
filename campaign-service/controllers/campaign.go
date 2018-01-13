package controllers

import (
	"log"
	// Standard packages
	"encoding/json"
	"fidelize/campaign-service/common"
	"fidelize/campaign-service/models"
	"fidelize/campaign-service/rules"
	"fmt"
	"net/http"

	//Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CampaignController struct represents the controller to perform get,post and delete
type CampaignController struct {
	session *mgo.Session
}

// NewCampaignController Gets a reference to CampaignController with a referenced mongo session
func NewCampaignController(s *mgo.Session) *CampaignController {
	return &CampaignController{s}
}

//GetAllCampaigns retrieves all campaigns
func (cc CampaignController) GetAllCampaigns(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Println("Getting all campaigns...")

	//Verify JWT
	//tokenIsValid, msgError := auth.ValidateToken(r)

	//if tokenIsValid {
	results := []models.Campaign{}
	// Find Anyone in mongoDB "CAMPAIGNS"
	if err := uc.session.DB(common.AppSettings.DBName).C("campaigns").Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	campaignsJSON, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", campaignsJSON)
	// } else {
	// 	//If Token was reject
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprint(w, msgError)
	// }

	log.Println("retriving all campaigns...")
}

// GetCampaign retrieves an individual campaign
func (cc CampaignController) GetCampaign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	//tokenIsValid, msgError := auth.ValidateToken(r)

	//if tokenIsValid {uc
	//Get parameter id
	id := p.ByName("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	//Get the verified ID
	idMongo := bson.ObjectIdHex(id)

	//New campaign from models
	u := models.Campaign{}

	// Fetch campaign
	if err := uc.session.DB(common.AppSettings.DBName).C("campaigns").FindId(idMongo).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content		uj, _ := json.Marshal(u) -Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
	// } else {
	// 	//If Token was reject		// Insert the campaign to the mongo
	uc.session.DB(common.AppSettings.DBName).C("campaign").Insert(u)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprint(w, msgError)
	// }
}

// CreateCampaign creates a new campaign
func (uc CampaignController) CreateCampaign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	//tokenIsValid, msgError := auth.ValidateToken(r)
	//if tokenIsValid {
	//New campaign from models
	u := models.Campaign{}
	// Populate the campaign data
	json.NewDecoder(r.Body).Decode(&u)
	// Add an ObjectId
	u.ID = bson.NewObjectId()
	validated, ValidationMsg := rules.ValidateInsetion(&u)

	if validated {
		// Insert the campaign to the mongo
		uc.session.DB(common.AppSettings.DBName).C("campaign").Insert(u)
		log.Println("Inserting Campaign ", u.Name)
		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
		/*} else {
			//If Token was reject
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, msgError)
		}*/
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ValidationMsg)
	}

}

// RemoveCampaign removes an existing campaign
func (uc CampaignController) RemoveCampaign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {
		// Get parameter id
		id := p.ByName("id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}

		// Get verified parameter id
		idMongo := bson.ObjectIdHex(id)

		// Remove campaign
		if err := uc.session.DB(common.AppSettings.DBName).C("campaigns").RemoveId(idMongo); err != nil {
			w.WriteHeader(404)
			return
		}

		// Write status
		w.WriteHeader(200)
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, msgError)
	}
}
