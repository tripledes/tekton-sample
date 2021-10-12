package database

import (
	"context"
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

func New(dbUrl string) (*DbConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	col := client.Database("quotedb").Collection("quotes")
	return &DbConn{Client: client, Col: col}, err
}

func (d *DbConn) Close() error {
	return d.Client.Disconnect(d.ctx)
}

func (d *DbConn) FindAll() ([]types.Quote, error) {
	var quotes []types.Quote
	cur, err := d.Col.Find(context.Background(), bson.D{})

	if err != nil {
		return quotes, err
	}

	if err := cur.All(context.Background(), &quotes); err != nil {
		return quotes, err
	}

	return quotes, nil
}
