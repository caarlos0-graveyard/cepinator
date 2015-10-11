package cep

import (
	"net/http"

	"github.com/caarlos0/cepinator/datastore"
	"github.com/labstack/echo"
)

func Index(ds datastore.Datastore) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		ceps, err := ds.LastUpdatedCeps(10)
		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, ceps)
	}
}
