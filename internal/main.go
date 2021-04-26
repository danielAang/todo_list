package internal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Constants
	Database *mongo.Database
}

func New() (*Config, error) {
	config := &Config{}
	constants, err := initViper()
	if err != nil {
		return config, err
	}
	config.Constants = constants
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if db, err := mongo.Connect(ctx, options.Client().ApplyURI(constants.Mongo.Url)); err != nil {
		return config, err
	} else {
		if err = db.Ping(ctx, nil); err != nil {
			log.Println("Unable to ping database", err)
			return config, err
		}
		config.Database = db.Database(config.Mongo.DbName)
		return config, err
	}
}
