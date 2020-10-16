package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ca17/teamsacs/config"
)

func GetMongodbClient(cfg config.MongodbConfig) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Url))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

