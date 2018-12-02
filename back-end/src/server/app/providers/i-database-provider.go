package providers

import "github.com/mongodb/mongo-go-driver/mongo"

type IDatabaseProvider interface {
	GetCollection(name string) (*mongo.Collection, error)
}
