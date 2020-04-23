package db

import (
	"context"
	"fmt"
	"github.com/benka-me/laruche-server/go-pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
	Bees   *mongo.Collection
	Hives  *mongo.Collection
}

func Init(cfg *config.Config) *DB {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.DbHost, cfg.DbPort)
	credential := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPWD,
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}
	return &DB{
		Client: client,
		Bees:   client.Database(config.LarsrvDatabase).Collection(config.BeesCollection),
		Hives:  client.Database(config.LarsrvDatabase).Collection(config.HivesCollection),
	}
}
