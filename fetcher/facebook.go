package fetcher

import (
	fb "github.com/huandu/facebook"
	"github.com/k0kubun/pp"

	"os"
)

type FacebookFetcher struct {
	pageId string
}

func (ff *FacebookFetcher) Fetch() (title string, body string, urls []string, err error) {
	res, _ := fb.Get("/"+ff.pageId+"/posts", fb.Params{
		"fields":       "message,full_picture,created_time,permalink_url,source",
		"access_token": os.Getenv("FACEBOOK_ACCESS_TOKEN"),
	})
	pp.Print(res)
	title = ""
	body = ""
	urls = []string{"hoge"}
	err = nil
	return
}
