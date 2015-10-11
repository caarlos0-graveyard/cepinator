package main

import (
	"log"

	"github.com/caarlos0/cepinator/config"
	"github.com/caarlos0/cepinator/datastore/database"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	defer db.Close()

	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/ping", pingHandler)
	e.Get("/ceps", cepIndex)
	e.Get("/ceps/:cep", cepSearch)

	log.Println("Running on port", cfg.Port)
	e.Run(":" + cfg.Port)
}
