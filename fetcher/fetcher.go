package fetcher

type FeedItem struct {
	Title         string
	Body          string
	ThumbnailPath string
	VideoPath     string
}

type Result struct {
	Name  string
	Items []FeedItem
	Error error
}

type Fetcher interface {
	Fetch() *Result
}
