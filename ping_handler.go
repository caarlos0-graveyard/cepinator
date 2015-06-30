package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func pingHandler(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
