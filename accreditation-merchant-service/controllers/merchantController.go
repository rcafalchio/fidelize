package controllers

import (
	"log"
	// Standard packages
	"encoding/json"
	"fidelize/accreditation-merchant-service/common"
	"fidelize/accreditation-merchant-service/models"
	"fidelize/accreditation-merchant-service/rules"
	"fmt"
	"net/http"

	//Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MerchantController struct represents the controller to perform get,post and delete
type MerchantController struct {
	session *mgo.Session
}

// NewMerchantController Gets a reference to MerchantController with a referenced mongo session
func NewMerchantController(s *mgo.Session) *MerchantController {
	return &MerchantController{s}
}

//GetAllMerchants retrieves all Merchants
func (uc MerchantController) GetAllMerchants(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Println("Getting all Merchants...")

	//Verify JWT
	//tokenIsValid, msgError := auth.ValidateToken(r)

	//if tokenIsValid {
	results := []models.Merchant{}
	// Find Anyone in mongoDB "MerchantS"
	if err := uc.session.DB(common.AppSettings.DBName).C("Merchants").Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	MerchantsJSON, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", MerchantsJSON)
	// } else {
	// 	//If Token was reject
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprint(w, msgError)
	// }

	log.Println("retriving all Merchants...")
}

// GetMerchant retrieves an individual Merchant
func (uc MerchantController) GetMerchant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

	//New Merchant from models
	u := models.Merchant{}

	// Fetch Merchant
	if err := uc.session.DB(common.AppSettings.DBName).C("Merchants").FindId(idMongo).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content		uj, _ := json.Marshal(u) -Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
	// } else {
	// 	//If Token was reject		// Insert the Merchant to the mongo
	uc.session.DB(common.AppSettings.DBName).C("Merchant").Insert(u)
	// 	w.WriteHeader(http.StatusUnautCPFhorized)
	// 	fmt.Fprint(w, msgError)
	// }
}

// CreateMerchant creates a new Merchant
func (uc MerchantController) CreateMerchant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	//tokenIsValid, msgError := auth.ValidateToken(r)
	//if tokenIsValid {
	//New Merchant from modelsnd
	m := models.Merchant{}
	// Populate the Merchant data
	json.NewDecoder(r.Body).Decode(&m)
	// Add an ObjectId
	m.ID = bson.NewObjectId()
	validated, ValidationMsg := rules.ValidateInsetion(&m)

	if validated {

		rules.FindSegment(&m)
		// Insert the Merchant to the mongo
		uc.session.DB(common.AppSettings.DBName).C("Merchant").Insert(m)
		log.Println("Inserting Merchant ", m.Name)
		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(m)
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

// RemoveMerchant removes an existing Merchant
func (uc MerchantController) RemoveMerchant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

		// Remove Merchant
		if err := uc.session.DB(common.AppSettings.DBName).C("Merchants").RemoveId(idMongo); err != nil {
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
