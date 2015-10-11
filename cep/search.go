package cep

import (
	"net/http"

	"github.com/caarlos0/cepinator/datastore"
	"github.com/labstack/echo"
)

func Search(ds datastore.Datastore) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		cep, err := ds.SearchCep(c.Param("cep"))
		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, cep)
	}
}
