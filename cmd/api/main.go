package main

import (
	"context"
	"erozen/mutants/internal/db"
	"erozen/mutants/internal/mutants"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("starting server...")

	mongoClient := newMongoClient()
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	db := db.New(mongoClient)
	service := mutants.NewService(db)
	handler := &MutantHandler{service: service}

	r := chi.NewRouter()
	r.Post("/mutant", handler.Mutant)
	r.Get("/stats", handler.Stats)
	http.ListenAndServe(":8080", r)
}

func newMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}
