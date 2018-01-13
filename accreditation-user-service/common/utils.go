package common

import (
	//Standard packages
	"encoding/json"
	"log"
	"os"

	//Third party packages
	"gopkg.in/mgo.v2"
)

type configuration struct {
	DBUser   string
	DBPws    string
	DBServer string
	DBName   string
	Server   string
}

//AppSettings represents all the config's like the database name, server , user and password
var AppSettings configuration

func loadSettings() {
	// //Getting the config file
	file, err := os.Open("settings-user-service.json")

	defer file.Close()

	// //Checking for error in open file
	if err != nil {
		log.Fatalf("We got a problem on JSON reading! \n%s", err)
	}

	// //Decoding the input file to json
	decoder := json.NewDecoder(file)

	// //Populating the struct that represents de config file
	err = decoder.Decode(&AppSettings)

	// //Checking for error
	if err != nil {
		log.Fatalf("We got a problem loadSettings() ! \n %s", err)
	}
}

// GetMongoSession creates a new mongo session and panics if connection error occurs
func GetMongoSession() *mgo.Session {

	log.Println("Connecting to ", AppSettings.DBServer)

	// Connect to mongo
	session, err := mgo.Dial(AppSettings.DBServer)
	//Check for error
	if err != nil {
		panic(err)
	}

	log.Println("Connected")
	// return de MongoSession
	return session
}
