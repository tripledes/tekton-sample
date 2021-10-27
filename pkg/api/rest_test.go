package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tripledes/web-quotes/fakedb"
	"github.com/tripledes/web-quotes/pkg/types"
)

func Test_getOneQuote(t *testing.T) {
	var q types.Quote
	db := fakedb.FakeDB{}
	app := NewWebApp(db)
	r := app.SetupServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quotes/one", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("Got HTTP code %d, expected 200", w.Code)
	}

	err := json.NewDecoder(w.Body).Decode(&q)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func Test_getAllQuotes(t *testing.T) {
	var qs []types.Quote
	db := fakedb.FakeDB{}
	app := NewWebApp(db)
	r := app.SetupServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quotes/all", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("Got HTTP code %d, expected 200", w.Code)
	}
	err := json.NewDecoder(w.Body).Decode(&qs)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}
