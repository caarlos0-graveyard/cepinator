package database

import (
	"database/sql"
	"log"

	"github.com/caarlos0/cepinator/datastore"
	"github.com/jmoiron/sqlx"
)

func Connect(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// NewDatastore returns a new Datastore
func NewDatastore(db *sql.DB) datastore.Datastore {
	dbx := sqlx.NewDb(db, "postgres")
	return struct {
		*Cepstore
	}{
		NewCepstore(dbx),
	}
}
