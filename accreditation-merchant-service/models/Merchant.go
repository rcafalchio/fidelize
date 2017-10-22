package models

//Third party packages
import "gopkg.in/mgo.v2/bson"

//Merchant model
type Merchant struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	CNPJ        string        `json:"cnpj" bson:"cnpj"`
	TradingName string        `json:"trading_price" bson:"trading_price"`
	Segment     int           `json:"segment" bson:"segment"`
	Mail        string        `json:"mail" bson:"mail"`
}
