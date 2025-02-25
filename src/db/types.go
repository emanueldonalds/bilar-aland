package db

type Listing struct {
	Url           string
	Title         string
	TitleTruncated    string
    Price         string
    CreatedAt     string
}

type ScrapeEvent struct {
	Date        string
}
