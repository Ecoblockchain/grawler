package fetcher

import (
	"github.com/labstack/echo"
	"net/http"
)

func FetchHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Fetch!")
}
