package models

//Third party packages

//Segment model
type Segment struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
