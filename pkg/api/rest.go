package api

import (
	"log"
	"net/http"

	"github.com/tripledes/web-quotes/pkg/database"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
	c  *gin.Context
	db *database.DbConn
}

func NewWebApp(db *database.DbConn) *WebApp {
	return &WebApp{}
}

func (w *WebApp) GetAllQuotes(c *gin.Context) {
	quotes, err := w.db.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	w.c.IndentedJSON(http.StatusOK, quotes)
}
