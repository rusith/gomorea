package providers

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"server/app/models"
	"server/app/providers"
)

type UserProvider struct {
	db providers.IDatabaseProvider
}

func NewUserProvider(db providers.IDatabaseProvider) providers.IUserProvider {
	return &UserProvider{ db }
}

func (u *UserProvider) GetByUsername(username string) (*models.User, error) {

	usersCollection, err := u.db.GetCollection("Users")
	if err != nil { return nil, err }
	result := bson.Raw{}
	a := usersCollection.FindOne(context.Background(), bson.D{{"Username", username}})
	a.Decode(&result)
	user := &models.User{}
	bson.Unmarshal([]byte(result), user)
	if err != nil { return nil, err }

	return user, nil
}