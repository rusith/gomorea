package account

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(email, password string) (bool, string){
	client, err := mongo.NewClient("mongodb://localhost:27017/server")
	if err != nil {
		return false, "An Error Occurred"
	}

	user := struct{
		Email string
		Password string
	}{}

	usersCollection := client.Database("server").Collection("users")
	err = usersCollection.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		return false, "Authorisation failed"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, "Authorisation failed"
	}

	return true, ""
}