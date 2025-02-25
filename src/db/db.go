package db

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/emanueldonalds/bilkoll/logger"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() *sql.DB {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPass := os.Getenv("MYSQL_PWD")

	if dbHost == "" {
		panic("MYSQL_HOST must be set.")
	}
	if dbPass == "" {
		panic("MYSQL_PWD must be set.")
	}

	connString := fmt.Sprintf("bilkoll:%s@tcp(%s:3306)/bilkoll", dbPass, dbHost)

	logger.Debug("Opening DB connection")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	logger.Debug("DB connection opened")

	db.SetConnMaxLifetime(180)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func GetListings(searchInput string, sqldb *sql.DB) []Listing {

	words := strings.Fields(searchInput)

	words = sanitize(words)
    words = expandHyphenatedWords(words)
    words = removeDashes(words)
    words = mercedes(words)

    search := strings.Join(words, " ")
    logger.Debug("Search query is " + search);

	params := []any{}
	var queryString = "SELECT " +
		"url, " +
		"title, " +
		"price " +
		"FROM listing "

	if search != "" {
		params = append(params, search)
		queryString = queryString + "WHERE MATCH(normalized_title) AGAINST(? IN BOOLEAN MODE) "
	}

    queryString = queryString + "ORDER BY normalized_title COLLATE 'utf8mb4_swedish_ci' ASC"

	logger.Debug("Get listings: Executing DB query")
	query, err := sqldb.Query(queryString, params...)

	if err != nil {
		panic(err.Error())
	}

	listings := []Listing{}

	for query.Next() {
		var rowListing Listing
		err := query.Scan(
			&rowListing.Url,
			&rowListing.Title,
			&rowListing.Price,
		)

		if err != nil {
			panic(err.Error())
		}

        // Create truncated title
        maxChars := 127;
        if (len(rowListing.Title) < maxChars) {
            rowListing.TitleTruncated = rowListing.Title
        } else {
            rowListing.TitleTruncated = rowListing.Title[:maxChars] + "..."
        }

		listings = append(listings, rowListing)
	}

	query.Close()
	logger.Debug("Get listings: Completed DB query")

	return listings
}

func GetLastScrape(sqldb *sql.DB) ScrapeEvent {
	query, qErr := sqldb.Query("SELECT date from scrape_event ORDER BY date DESC LIMIT 1")

	if qErr != nil {
		panic(qErr.Error())
	}

	query.Next()

	var scrapeEvent ScrapeEvent

	logger.Debug("Last scrape: Executing DB query")
	sErr := query.Scan(&scrapeEvent.Date)

	query.Close()
	logger.Debug("Last scrape: Completed DB query")

	if sErr != nil {
		panic(sErr.Error())
	}

	return scrapeEvent
}

func sanitize(words []string) []string {

    var sanitized = []string{}
	for _, word := range words {
        switch word {
            case "-", "+", "*", "/", "\\", "=":
            default:
                re := regexp.MustCompile("[^A-Za-z0-9\\- ]")
                sanitized = append(sanitized, re.ReplaceAllString(word, ""));
        }
    }
	return sanitized
}


// In case of hyphenated word, add the individual words as separate search terms, 
// e.g. ["Mercedes-Benz"] -> ["Mercedes-Benz", "Mercedes", "Benz"]
func expandHyphenatedWords(words []string) []string {

    logger.Debug("Input:")
    for _, word := range words {
        logger.Debug(word);
    }
    
    var result = []string{}
	for _, word := range words {
        // Skip if word starts or ends with '-'
        if word[0:1] == "-" || word[len(word)-1:] == "-" {
            result = append(result, word)
        } else {
            parts := strings.Split(word, "-");
            result = append(result, word);
            if len(parts) > 1 {
                result = append(result, parts...);
            }
        }
	}
    logger.Debug("Output:")
    for _, word := range result {
        logger.Debug(word);
    }
    return result
}

// Special handling for mercedes to make query for just "Mercedes" also match "MercedesBenz" and "Mercedes-Benz"
func mercedes(words []string) []string {
    var res = []string{}
    for _,word := range(words) {
        if strings.ToLower(word) == "mercedes" {
            res = append(res, "MercedesBenz")
        }
        res = append(res, word)
    }
    return res
}

func removeDashes(words []string) []string {

    var withoutDashes = []string{}
	for _, word := range words {
        res := word[0:1] + strings.Replace(word[1:], "-", "", -1)
        if len(res) == 1 {
            res = strings.Replace(res, "-", "", -1)
        }
        withoutDashes = append(withoutDashes, res)
	}
    return withoutDashes
}
