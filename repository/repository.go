package repository

import (
	"context"

	"github.com/engine/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryI interface {
	Get(ctx context.Context) (models.Response, error)
}

type Repository struct {
	db *mongo.Database
}

func NewRepository(mongo *mongo.Database) *Repository {
	return &Repository{
		db: mongo,
	}
}

func (r *Repository) Get(ctx context.Context) (models.Response, error) {
	collections, err := r.db.Collection("restaurants").Find(ctx, bson.M{})
	if err != nil {
		return models.Response{
			Message:     "error",
			Restaurants: nil,
		}, err
	}

	var restaurants []models.Restaurant
	for collections.Next(ctx) {
		var restaurant models.Restaurant
		err := collections.Decode(&restaurant)
		if err != nil {
			return models.Response{
				Message:     "error",
				Restaurants: nil,
			}, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return models.Response{
		Message:     "success",
		Restaurants: restaurants,
	}, nil
}
