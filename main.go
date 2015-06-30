package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var db gorm.DB

func main() {
	db = NewDbConnectionPool()
	defer db.Close()

	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/ping", pingHandler)
	e.Get("/ceps", cepIndex)
	e.Get("/ceps/:cep", cepSearch)

	port := Getenv("PORT", "3000")
	log.Println("Running on port", port)
	e.Run(":" + port)
}
