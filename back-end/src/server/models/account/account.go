package account

import "github.com/mongodb/mongo-go-driver/mongo"

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func Login(email, password string) (map[string]interface{}) {
	client, err := mongo.NewClient("mongodb://foo:bar@localhost:27017")
}