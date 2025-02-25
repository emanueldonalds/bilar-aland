package web

import (
	"database/sql"
	"net/http"

	"github.com/emanueldonalds/bilkoll/db"
)

func IndexHandler(sqldb *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		listings := db.GetListings(search, sqldb)
		lastScrape := db.GetLastScrape(sqldb)
		index := Index(listings, lastScrape)
		index.Render(r.Context(), w)
	})

}

func FilterHandler(sqldb *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		listings := db.GetListings(search, sqldb)
		lastScrape := db.GetLastScrape(sqldb)
		index := Listings(listings, lastScrape)
		index.Render(r.Context(), w)
	})
}
