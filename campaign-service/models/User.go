package models

import "gopkg.in/mgo.v2/bson"

//User model
type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Mail   string        `json:"mail" bson:"mail"`
	CPF    string        `json:"cpf" bson:"cpf"`
	Gender string        `json:"gender" bson:"gender"`
}
