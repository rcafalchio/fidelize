package models

//Third party packages
import "gopkg.in/mgo.v2/bson"

//Widget model
type Merchant struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Color     string        `json:"color" bson:"color"`
	Price     string        `json:"price" bson:"price"`
	Inventory int           `json:"inventory" bson:"inventory"`
	Melts     bool          `json:"melts" bson:"melts"`
}
