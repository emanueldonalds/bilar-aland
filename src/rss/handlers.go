package rss

import (
	"bytes"
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"os"
	"text/template"

	"github.com/emanueldonalds/bilkoll/db"
	"github.com/emanueldonalds/bilkoll/formatters"
)

var rssTemplate = readFile("rss/template.xml")

func RssHandler(sqldb *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		listings := db.GetListings(search, "title", "desc", sqldb)
		lastScrape := db.GetLastScrape(sqldb)

		items := []Item{}
		for i := 0; i < len(listings); i++ {
			listing := listings[i]

			items = append(items, Item{
				Title: listing.Title,
				Link:  listing.Url,
				Description: fmt.Sprintf(
					html.EscapeString("<ul><li>Price: %s</li></ul>"),
					formatters.FormatPrice(listing.Price),
				),
				PubDate: formatters.FormatDateTimeRfc822(listing.CreatedAt),
				Id:      listing.Url,
			})
		}

		data := Feed{
			Title:       "Bilar till salu på Åland",
			Description: "Bilar till salu på Åland RSS feed",
			PubDate:     formatters.FormatDateTimeRfc822(lastScrape.Date),
			WebMaster:   "husax@protonmail.com",
			Items:       items,
		}

		tmpl, err := template.New("feed").Parse(rssTemplate)
		if err != nil {
			panic(err)
		}

		var res bytes.Buffer
		tmpl.Execute(&res, data)

		w.Header().Set("Content-Type", "application/rss+xml")
		w.Write([]byte(res.Bytes()))
	})
}

func readFile(filename string) string {
	fmt.Println("Loading " + filename)
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("Error reading file at " + pwd + "/" + filename)
		panic(err)
	}
	content := string(fileBytes)
	return content
}
