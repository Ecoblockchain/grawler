package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Root!")
}

func registrationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Register!")
}

func fetchHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Fetch!")
}

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.POST("/register", registrationHandler)
	e.POST("/fetch/:id", fetchHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
