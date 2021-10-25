package database

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/tripledes/web-quotes/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	Close() error
	FindAll() ([]types.Quote, error)
	FindOne() (types.Quote, error)
}

func NewDB(dbUrl string) (DB, error) {
	re := regexp.MustCompilePOSIX(`^mongodb://.*`)
	if re.MatchString(dbUrl) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
		col := client.Database("quotesdb").Collection("quotes")
		return &MongoClient{Client: client, Col: col}, err
	}

	return nil, fmt.Errorf("DB for %s is not supported", dbUrl)
}
