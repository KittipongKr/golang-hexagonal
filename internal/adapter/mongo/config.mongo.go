package mongo

import (
	"context"
	"fmt"

	envCfgs "csat-servay/configs/env"
	m "csat-servay/internal/adapter/mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfgs *envCfgs.MongoConfig) (*m.MongoCollections, error) {

	dsn := fmt.Sprintf("mongodb://%s:%s/",
		cfgs.Host,
		cfgs.Port,
	)

	opts := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfgs.Database)

	collections := &m.MongoCollections{
		Users: db.Collection("users"),
	}

	return collections, nil
}
