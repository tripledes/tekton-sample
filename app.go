package main

import (
	"log"
	"os"

	"github.com/tripledes/web-quotes/pkg/database"

	"github.com/tripledes/web-quotes/pkg/api"

	"github.com/gin-gonic/gin"
)

func main() {

	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		dbUrl = "mongodb://localhost:27017"
	}

	db, err := database.New(dbUrl)
	app := api.NewWebApp(db)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.GET("/quotes/all", app.GetAllQuotes)

	router.Run(":8080")

}
