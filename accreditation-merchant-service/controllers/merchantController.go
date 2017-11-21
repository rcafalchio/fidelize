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

	results := []models.Merchant{}

	if err := uc.session.DB(common.AppSettings.DBName).C("Merchant").Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	MerchantsJSON, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	//Get the verified ID
	idMongo := bson.ObjectIdHex(id)
	//New Merchant from models
	m := models.Merchant{}

	// Fetch Merchant
	if err := uc.session.DB(common.AppSettings.DBName).C("Merchant").FindId(idMongo).One(&m); err != nil {
		if fmt.Sprint(err) == "not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Fatal("Erro ao buscar o usuário", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	println("BUSCOU = ", m.ID)

	uj, _ := json.Marshal(m)
	// Write content-type, statuscode, payload
	w.Header().Set("Content		uj, _ := json.Marshal(u) -Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
	// } else {
	// 	//If Token was reject
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
		m.Active = true
		// Insert the Merchant to the mongo
		uc.session.DB(common.AppSettings.DBName).C("Merchant").Insert(m)
		log.Println("Inserting Merchant ", m.Name)
		// Marshal provided interface into JSON structure
		mj, _ := json.Marshal(m)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", mj)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ValidationMsg)
	}

}

// InactivateMerchant - Inactivate an existing Merchant
func (uc MerchantController) InactivateMerchant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Get parameter id
	id := p.ByName("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	// Get verified parameter id
	idMongo := bson.ObjectIdHex(id)
	change := bson.M{"$set": bson.M{"active": false}}

	// Remove Merchant
	if err := uc.session.DB(common.AppSettings.DBName).C("Merchant").UpdateId(idMongo, change); err != nil {
		if fmt.Sprint(err) == "not found" {
			println("TESTE RICARDO - Não encontrou!!!")
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Fatal("Erro ao desativar o usuário", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Write status
	w.WriteHeader(200)
}
