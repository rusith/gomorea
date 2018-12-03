package providers

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"server/app/models"
	"server/app/providers"
)

type UserProvider struct {
	db providers.IDatabaseProvider
}

func NewUserProvider(db providers.IDatabaseProvider) providers.IUserProvider {
	return &UserProvider{ db }
}

func (u *UserProvider) GetCollection() (*mgo.Collection, error){
	collection, err := u.db.GetCollection("Users")
	return collection, err
}

func (u *UserProvider) GetByUsername(username string) (*models.User, error) {

	col, err := u.GetCollection()
	if err != nil { return nil, err }
	user := &models.User{}
	err = col.Find(bson.M{"username": username}).One(user)
	if err != nil { return nil, err }

	return user, nil
}

func (u *UserProvider) Add(user *models.User) (*models.User, error) {
	col, _ := u.GetCollection()
	user.ID = bson.NewObjectId()
	i := col.Insert(user)
	return user, i
}