package database

import (
	"github.com/caarlos0/cepinator/datastore/model"
	"github.com/jmoiron/sqlx"
)

// Cepstore store ceps in database
type Cepstore struct {
	*sqlx.DB
}

// NewCepstore datastore
func NewCepstore(db *sqlx.DB) *Cepstore {
	return &Cepstore{db}
}

func (db *Cepstore) LastUpdatedCeps(amount int) ([]model.CEP, error) {
	var ceps []model.CEP
	return ceps, db.Select(
		&ceps, "select * from ceps order by updated_at desc limit $1", amount,
	)
}

func (db *Cepstore) SearchCep(query string) (model.CEP, error) {
	var cep model.CEP
	return cep, db.Get(
		&cep, "select * from ceps where value = $1 limit 1", query,
	)
}
