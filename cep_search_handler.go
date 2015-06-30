package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func cepSearch(c *echo.Context) error {
	result := db.Where("value = ?", c.Param("cep")).First(&Cep{})
	if result.RowsAffected == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, result.Value)
}
