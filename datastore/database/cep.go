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

var insertCepStm = `
insert into ceps(
	city, state, uf, logradouro, neighborhood, address, complement, value
) values (
	:city, :state, :uf, :logradouro, :neighborhood, :address, :complement, :value
) returning id
`

func (db *Cepstore) CreateCep(cep model.CEP) (model.CEP, error) {
	rows, err := db.NamedQuery(insertCepStm, cep)
	if err != nil {
		return model.CEP{}, err
	}
	defer rows.Close()
	var id int64
	rows.Next()
	if err := rows.Scan(&id); err != nil {
		return model.CEP{}, err
	}
	return db.FindCepByID(id)
}

var updateCepStm = `
update ceps
where id = :id
set
	city = :city,
	state = :state,
	uf = :uf,
	logradouro = :logradouro,
	neighborhood = :neighborhood,
	address = :address,
	complement = :complement,
	value = :value
`

func (db *Cepstore) UpdateCep(cep model.CEP) (model.CEP, error) {
	_, err := db.NamedExec(insertCepStm, cep)
	if err != nil {
		return model.CEP{}, err
	}
	return db.FindCepByID(cep.ID)
}

func (db *Cepstore) FindCepByID(id int64) (model.CEP, error) {
	var result model.CEP
	return result, db.Get(&result, "select * from ceps where id = $1", id)
}
