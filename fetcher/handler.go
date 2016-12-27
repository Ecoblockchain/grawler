package fetcher

import (
	"github.com/labstack/echo"
	"net/http"
)

func FetchHandler(c echo.Context) error {
	fbFetcher := &FacebookFetcher{pageId: 538744468}
	a, b, c, d := fbFetcher.Fetch()
	return c.String(http.StatusOK, "Fetch!")
}
