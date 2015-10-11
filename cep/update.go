package cep

import (
	"net/http"

	"github.com/caarlos0/cepinator/datastore"
	"github.com/caarlos0/cepinator/datastore/model"
	"github.com/labstack/echo"
)

func Update(ds datastore.Datastore) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var cep model.CEP
		if err := c.Bind(&cep); err != nil {
			return err
		}
		cep.Value = c.Param("cep")
		updatedCep, err := ds.UpdateCep(cep)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, updatedCep)
	}
}
