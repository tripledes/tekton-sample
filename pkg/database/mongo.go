package database

import (
	"context"
	"log"

	"github.com/tripledes/web-quotes/pkg/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
	Col    *mongo.Collection
	ctx    context.Context
}

func (m *MongoClient) Close() error {
	return m.Client.Disconnect(m.ctx)
}

func (m *MongoClient) FindAll() ([]types.Quote, error) {
	var quotes []types.Quote

	findOptions := options.Find()
	findOptions.SetLimit(5)
	cur, err := m.Col.Find(m.ctx, bson.D{})
	if err != nil {
		log.Println("failed to find all quotes")
		return quotes, err
	}

	defer cur.Close(m.ctx)

	if err := cur.All(context.Background(), &quotes); err != nil {
		log.Printf("error decoding quotes: %v\n", err.Error())
		return quotes, err
	}

	log.Printf("found %d quotes", len(quotes))
	return quotes, nil
}

func (m *MongoClient) FindOne() (types.Quote, error) {
	var q types.Quote
	err := m.Col.FindOne(m.ctx, bson.M{}).Decode(&q)
	if err != nil {
		log.Printf("error finding one: %v", err)
		return q, err
	}
	log.Printf("One quote: %v", q)
	return q, err
}
