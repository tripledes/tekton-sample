package database

import (
	"context"
	"log"
	"time"

	"github.com/tripledes/web-quotes/pkg/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConn struct {
	Client *mongo.Client
	Col    *mongo.Collection
	ctx    context.Context
}

func NewDB(dbUrl string) (*DbConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	col := client.Database("quotesdb").Collection("quotes")
	return &DbConn{Client: client, Col: col}, err
}

func (d *DbConn) Close() error {
	return d.Client.Disconnect(d.ctx)
}

func (d *DbConn) FindAll() ([]types.Quote, error) {
	var quotes []types.Quote

	findOptions := options.Find()
	findOptions.SetLimit(5)
	cur, err := d.Col.Find(d.ctx, bson.D{})
	if err != nil {
		log.Println("failed to find all quotes")
		return quotes, err
	}

	defer cur.Close(d.ctx)

	if err := cur.All(context.Background(), &quotes); err != nil {
		log.Printf("error decoding quotes: %v\n", err.Error())
		return quotes, err
	}

	log.Printf("found %d quotes", len(quotes))
	return quotes, nil
}

func (d *DbConn) FindOne() (types.Quote, error) {
	var q types.Quote
	err := d.Col.FindOne(d.ctx, bson.M{}).Decode(&q)
	if err != nil {
		log.Printf("error finding one: %v", err)
		return q, err
	}
	log.Printf("One quote: %v", q)
	return q, err
}
