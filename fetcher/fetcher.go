package fetcher

type Fetcher interface {
	Fetch() (title string, body string, urls []string, err error)
}
