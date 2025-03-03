package db

type Listing struct {
	Url           string
	Title         string
	TitleTruncated    string
    Price         int32
    CreatedAt     string
}

type ScrapeEvent struct {
	Date        string
}
