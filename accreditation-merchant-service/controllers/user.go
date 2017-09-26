package controllers

import (
	// Standard packages
	"encoding/json"
	"fmt"
	"net/http"

	//Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserController struct represents the controller to perform get,post and delete
type UserController struct {
	session *mgo.Session
}

// NewUserController Gets a reference to UserController with a referenced mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//GetAllUsers retrieves all users
func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {
		results := []models.User{}
		// Find Anyone in mongoDB "USERS"
		if err := uc.session.DB(common.AppSettings.DBName).C("users").Find(nil).All(&results); err != nil {
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

// GetUser retrieves an individual user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
		u := models.User{}

		// Fetch user
		if err := uc.session.DB(common.AppSettings.DBName).C("users").FindId(idMongo).One(&u); err != nil {
			w.WriteHeader(404)
			return
		}

		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content		uj, _ := json.Marshal(u)
		-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	} else {
		//If Token was reject
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, msgError)
	}
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//Verify JWT
	tokenIsValid, msgError := auth.ValidateToken(r)

	if tokenIsValid {
		//New user from models
		u := models.User{}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(&u)

		// Add an ObjectId
		u.Id = bson.NewObjectId()

		// Insert the user to the mongo
		uc.session.DB(common.AppSettings.DBName).C("users").Insert(u)

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

// RemoveUser removes an existing user
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
		if err := uc.session.DB(common.AppSettings.DBName).C("users").RemoveId(idMongo); err != nil {
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
