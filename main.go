package main

import (
	"log"

	"github.com/caarlos0/cepinator/cep"
	"github.com/caarlos0/cepinator/config"
	"github.com/caarlos0/cepinator/datastore/database"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	defer db.Close()
	ds := database.NewDatastore(db)

	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	e.Get("/ping", ping.Index())
	e.Get("/ceps", cep.Index(ds))
	e.Get("/ceps/:cep", cep.Search(ds))

	log.Println("Running on port", cfg.Port)
	e.Run(":" + cfg.Port)
}
