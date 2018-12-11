package providers

import (
	"github.com/globalsign/mgo"
	"server/app/config"
	"server/app/providers"
)

type DatabaseProvider struct {
	config config.IAppConfig
}

func NewDatabaseProvider(config config.IAppConfig) providers.IDatabaseProvider {
	return &DatabaseProvider{
		config,
	}
}


func (d *DatabaseProvider) GetCollection(name string) (*mgo.Collection, error) {

	session, err := mgo.Dial(d.config.DatabaseUrl())
	c := session.DB("animus").C(name)
	if err != nil {
		return nil, err
	}

	return c, nil
}