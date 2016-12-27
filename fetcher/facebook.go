package fetcher

import (
	fb "github.com/huandu/facebook"

	"github.com/k0kubun/pp"

	"os"
)

// define a facebook feed object.
type FacebookFeed struct {
	// mind the "," before "required".
	Data []*FacebookFeedData `facebook:"data"` // use customized field name "from".
}

type FacebookFeedData struct {
	Message, FullPicture, Source, CreatedTime string
}

// create a feed object direct from graph api result.
var feed FacebookFeed

type FacebookFetcher struct {
	pageId string
}

func (ff *FacebookFetcher) Fetch() (result Result) {
	res, _ := fb.Get("/"+ff.pageId+"/posts", fb.Params{
		"fields":       "message,full_picture,created_time,permalink_url,source",
		"access_token": os.Getenv("FACEBOOK_ACCESS_TOKEN"),
	})

	res.Decode(&feed)
	pp.Print(feed)

	var feedItems []FeedItem
	for _, data := range feed.Data {
		data = data
		if data.Message == "" {
			continue
		}
		feedItems = append(feedItems, FeedItem{Title: "", Body: data.Message, ThumbnailPath: data.FullPicture, VideoPath: data.Source, CreatedAt: data.CreatedTime})
	}

	result = Result{Name: "", Items: feedItems, Error: nil}
	return result
}
