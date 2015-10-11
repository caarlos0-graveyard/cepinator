package ping

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index() func(c *echo.Context) error {
	return func(c *echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
}
