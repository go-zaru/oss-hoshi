package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-zaru/oss-hoshi/packages/nioh/pkgs/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConfig struct {
	ConnectionURL string `mapstructure:""`
}

var cfg mongoConfig

const key = "mongodb"

// CreateMongoDriver ...
func CreateMongoDriver(provider *config.CfgProvider) *mongo.Client {
	provider.UnmarshalExactKey(key, &cfg)
	driver := connect(cfg.ConnectionURL)
	return driver
}

func connect(url string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	driver, er := mongo.Connect(
		ctx, options.Client().ApplyURI(url),
	)

	if er != nil {
		panic(fmt.Errorf("Could not connect to mongo: %s", er))
	}

	return driver
}
