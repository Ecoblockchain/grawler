package fetcher

import (
	"github.com/huandu/facebook"
	"github.com/k0kubun/pp"
)

type FacebookFetcher struct {
	pageId string
}

func (ff *FacebookFetcher) Fetch() (title string, body string, urls []string, err error) {
	res, _ := fb.Get("/538744468", fb.Params{
		"fields":       "first_name",
		"access_token": "a-valid-access-token",
	})
	pp.Print(res)
	return nil, nil, nil, nil
}
