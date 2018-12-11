package models

import "github.com/globalsign/mgo/bson"

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}
