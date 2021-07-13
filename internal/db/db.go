package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	client *mongo.Client
}

func New(client *mongo.Client) *DB {
	return &DB{client: client}
}

func (d *DB) Save(ctx context.Context, test *Test) error {
	result, err := d.client.Database("mutants").Collection("tests").InsertOne(ctx, test)
	if err != nil {
		return err
	}

	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		test.ID = insertedID.Hex()
	}

	return nil
}

func (d *DB) Count(ctx context.Context, isMutant bool) (int64, error) {
	filter := bson.M{"mutant": isMutant}
	return d.client.Database("mutants").Collection("tests").CountDocuments(ctx, filter)
}
