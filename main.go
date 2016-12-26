package main

import (
	"github.com/labstack/echo"
	fetcher "github.com/timakin/grawler/fetcher"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Root!")
}

func registrationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Register!")
}

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.POST("/register", registrationHandler)
	e.POST("/fetch/:id", fetcher.FetchHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
