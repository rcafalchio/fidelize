package models

//Third party packages
import "gopkg.in/mgo.v2/bson"

//Merchant model
type Merchant struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Mail        string        `json:"mail" bson:"mail"`
	Password    string        `json:"password" bson:"password"`
	TradingName string        `json:"trading_name" bson:"trading_name"`
	CNPJ        int           `json:"cnpj" bson:"cnpj"`
	Segment     Segment       `json:"segment" bson:"segment"`
	Active      bool          `json:"active" bson:"active"`
}
