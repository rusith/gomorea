package models

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type User struct {
	Id objectid.ObjectID `json:"id"`
	Username string `json:"username",bson:"Username,omitempty"`
	Password string `json:"password",bson:"Password,omitempty"`
}
