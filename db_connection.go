package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const defaultDbUrl = "postgres://localhost:5432/cepinator?sslmode=disable"

func NewDbConnectionPool() gorm.DB {
	db, err := gorm.Open("postgres", Getenv("DATABASE_URL", defaultDbUrl))
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Cep{})
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}
