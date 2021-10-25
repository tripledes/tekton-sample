package api

import (
	"log"
	"net/http"

	"github.com/tripledes/web-quotes/pkg/database"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
	db *database.DbConn
}

func NewWebApp(db *database.DbConn) *WebApp {
	return &WebApp{db: db}
}

func (w *WebApp) GetAllQuotes(c *gin.Context) {
	quotes, err := w.db.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fetched quotes: %v", quotes)
	c.IndentedJSON(http.StatusOK, quotes)
}

func (w *WebApp) GetOneQuote(c *gin.Context) {
	quote, err := w.db.FindOne()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fetched quotes: %v", quote)
	c.IndentedJSON(http.StatusOK, quote)
}
