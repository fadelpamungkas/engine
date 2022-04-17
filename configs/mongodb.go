package configs

import (
	"context"
	"time"

	"github.com/engine/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(c models.DatabaseConfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(c.URI[1])
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(c.DBname), err
}
