package fetcher

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
	Crawl()
}
