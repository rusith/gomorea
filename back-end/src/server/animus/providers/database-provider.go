package providers

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"server/app/config"
	"server/app/providers"
	"time"
)

type DatabaseProvider struct {
	config config.IAppConfig
}

func NewDatabaseProvider(config config.IAppConfig) providers.IDatabaseProvider {
	return &DatabaseProvider{
		config,
	}
}


func (d *DatabaseProvider) GetCollection(name string) (*mongo.Collection, error) {
	fmt.Println(d.config.DatabaseUrl())
	client, err := mongo.NewClient(d.config.DatabaseUrl())
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("animus").Collection(name), nil
}