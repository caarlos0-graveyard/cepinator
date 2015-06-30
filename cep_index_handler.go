package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func cepIndex(c *echo.Context) error {
	result := db.Order("updated_at desc").Limit(10).Find(&[]Cep{})
	if result.RowsAffected == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, result.Value)
}
