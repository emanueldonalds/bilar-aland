package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/emanueldonalds/bilkoll/db"
	"github.com/emanueldonalds/bilkoll/rss"
	"github.com/emanueldonalds/bilkoll/web"
)

// Same etag for all files, generates a new every time server restarts
var etag string = "W/\"" + fmt.Sprint(time.Now().UTC().Unix()) + "\""

func main() {
	assetsDir := "./assets"

	info, err := os.Stat(assetsDir)

	if err != nil {
		panic("Could not stat assets directory. Make sure assets dir is in the working directory.")
	}
	if info.Mode().Perm()&0444 != 0444 {
		panic("Missing permissions to read assets")
	}

	mux := http.NewServeMux()
	db := db.GetDb()

	mux.Handle("/", web.IndexHandler(db))
	mux.Handle("/filter", web.FilterHandler(db))
	mux.Handle("/rss", rss.RssHandler(db))
	mux.Handle("/assets/", cacheControl(http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir)))))

	fmt.Println("Listening on :4942")
	log.Fatal(http.ListenAndServe(":4942", mux))
}

func cacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ifNoneMatch := r.Header.Get("If-None-Match")

		if ifNoneMatch == etag {
			w.WriteHeader(304)
			return
		}

		w.Header().Set("ETag", etag)
		w.Header().Set("Cache-Control", "no-cache, must-revalidate")
		h.ServeHTTP(w, r)
	})
}
