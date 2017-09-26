package controllers

import (
	// Standard packages
	"encoding/json"
	"fmt"
	"net/http"

	// My packages
	"github.com/pbribeiro/widget-go-spa-api/common"
	"github.com/pbribeiro/widget-go-spa-api/models"

	//Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// WidgetController struct represents the controller to perform get,post and delete
type WidgetController struct {
	session *mgo.Session
}

// NewWidgetController Gets a reference to WidgetController with a referenced mongo session
func NewWidgetController(s *mgo.Session) *WidgetController {
	return &WidgetController{s}
}

//GetAllWidgets retrieves all widgets
func (wc WidgetController) GetAllWidgets(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {

		results := []models.Widget{}

		// Find Anyone in mongoDB "WIDGETS"
		if err := wc.session.DB(common.AppSettings.DBName).C("widgets").Find(nil).All(&results); err != nil {
			w.WriteHeader(404)
			return
		}

		usersJSON, _ := json.Marshal(results)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", usersJSON)
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, msgError)
	}
}

// GetWidget retrieves an individual widget
func (wc WidgetController) GetWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {
		//Get parameter id
		id := p.ByName("id")

		// Verify id is ObjectId
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}

		//Get the verified ID
		idMongo := bson.ObjectIdHex(id)

		//New user from models
		u := models.Widget{}

		// Fetch widget
		if err := wc.session.DB(common.AppSettings.DBName).C("widgets").FindId(idMongo).One(&u); err != nil {
			w.WriteHeader(404)
			return
		}

		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, msgError)
	}
}

// CreateWidget creates a new widget
func (wc WidgetController) CreateWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {
		//New user from models
		u := models.Widget{}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(&u)

		// Add an ObjectId
		u.Id = bson.NewObjectId()

		// Insert the user to the mongo
		wc.session.DB(common.AppSettings.DBName).C("widgets").Insert(u)

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, msgError)
	}
}

// RemoveWidget removes an existing widget
func (wc WidgetController) RemoveWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

		// Remove user
		if err := wc.session.DB(common.AppSettings.DBName).C("widgets").RemoveId(idMongo); err != nil {
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

// UpdateWidget updates an existing widget
func (wc WidgetController) UpdateWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

		//New user from models
		u := models.Widget{}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(&u)

		// Remove user
		if err := wc.session.DB(common.AppSettings.DBName).C("widgets").UpdateId(idMongo, u); err != nil {
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
