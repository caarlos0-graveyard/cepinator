package database
import (
	"github.com/jmoiron/sqlx"
	"github.com/caarlos0/cepinator/datastore"
	"github.com/caarlos0/cepinator/datastore/model"
)

// Cepstore store ceps in database
type Cepstore struct {
	*sqlx.DB
}

// NewCepstore datastore
func NewCepstore(db *sqlx.DB) *datastore.Cepstore {
	return &Cepstore{db}
}

func (db *Cepstore) LastUpdatedCeps(amount int) ([]model.CEP, error) {
	var ceps []model.CEP
	return ceps, db.Select(
		&ceps, "select * from ceps order by updated_at desc limit $1", amount,
	)
}

func (db *Cepstore) SearchCep(cep string) (model.CEP, error) {
	var cep model.CEP
	return cep, db.Get(&cep, "select * from ceps where value = $1 limit 1", cep)
}
