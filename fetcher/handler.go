package fetcher

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/k0kubun/pp"
)

func FetchHandler(c echo.Context) error {
	pp.Print(c.QueryParam("page_id"))
	fbFetcher := &FacebookFetcher{pageId: c.QueryParam("page_id")}
	a, b, arraa, d := fbFetcher.Fetch()
	pp.Print(a)
	pp.Print(b)
	pp.Print(arraa)
	pp.Print(d)
	return c.String(http.StatusOK, "Fetch!")
}
