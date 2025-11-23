package mongo

import (
	"fmt"
	"time"

	envCfgs "csat-servay/configs/env"
	m "csat-servay/internal/adapter/mongo/model"
	r "csat-servay/internal/adapter/mongo/repository"
	p "csat-servay/internal/core/port"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/v2/mongo/otelmongo"
	"go.opentelemetry.io/otel/trace"
)

func Connect(cfgs *envCfgs.MongoConfig, tp trace.TracerProvider) (*m.MongoCollections, error) {
	dsn := fmt.Sprintf("mongodb://%s:%s/",
		cfgs.Host,
		cfgs.Port,
	)
	// opts := options.Client().ApplyURI(dsn)
	opts := options.Client()
	opts.ApplyURI(dsn)
	// Set timeout
	opts.SetConnectTimeout(10 * time.Second)
	opts.SetTimeout(10 * time.Second)
	// Set tracing monitor
	if tp != nil {
		opts.SetMonitor(
			otelmongo.NewMonitor(otelmongo.WithTracerProvider(tp)),
		)
	}

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfgs.Database)

	collections := &m.MongoCollections{
		Users: db.Collection("users"),
	}

	return collections, nil
}

type Adaptor struct {
	PingRepo p.PingRepo
}

func SetAdaptor(collections *m.MongoCollections) Adaptor {
	return Adaptor{
		PingRepo: r.NewPingRepo(),
	}
}
