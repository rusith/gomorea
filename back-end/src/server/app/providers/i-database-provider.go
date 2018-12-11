package providers

import "github.com/globalsign/mgo"

type IDatabaseProvider interface {
	GetCollection(name string) (*mgo.Collection, error)
}
