package web

import (
	"database/sql"
	"net/http"

	"github.com/emanueldonalds/bilkoll/db"
)

func IndexHandler(sqldb *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		sortBy := r.URL.Query().Get("sort")
		sortOrder := r.URL.Query().Get("order")
		listings := db.GetListings(search, sortBy, sortOrder, sqldb)
		lastScrape := db.GetLastScrape(sqldb)
		index := Index(listings, lastScrape, search, sortBy, sortOrder)
		index.Render(r.Context(), w)
	})

}
