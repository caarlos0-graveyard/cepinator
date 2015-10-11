package datastore

import "github.com/caarlos0/cepinator/datastore/model"

type Cepstore interface {
	LastUpdatedCeps(amount int) ([]model.CEP, error)
	SearchCep(cep string) (model.CEP, error)
	CreateCep(cep model.CEP) (model.CEP, error)
	UpdateCep(cep model.CEP) (model.CEP, error)
}
