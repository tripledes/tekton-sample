package api

import (
	"log"
	"net/http"

	"github.com/tripledes/web-quotes/pkg/database"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
	db database.DB
}

func NewWebApp(db database.DB) *WebApp {
	return &WebApp{db: db}
}

func (w *WebApp) SetupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/quotes/all", w.getAllQuotes)
	router.GET("/quotes/one", w.getOneQuote)
	return router
}

func (w *WebApp) getAllQuotes(c *gin.Context) {
	quotes, err := w.db.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fetched quotes: %v", quotes)
	c.IndentedJSON(http.StatusOK, quotes)
}

func (w *WebApp) getOneQuote(c *gin.Context) {
	quote, err := w.db.FindOne()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fetched quotes: %v", quote)
	c.IndentedJSON(http.StatusOK, quote)
}
