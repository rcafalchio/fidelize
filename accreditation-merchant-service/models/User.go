package models

import "gopkg.in/mgo.v2/bson"

//User model
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Gravatar string        `json:"gravatar" bson:"gravatar"`
}
