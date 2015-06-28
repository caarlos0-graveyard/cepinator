package main

import (
	"github.com/jinzhu/gorm"
)

type Cep struct {
	gorm.Model
	City         string `sql:"not null;index:idx_city_value"`
	State        string `sql:"not null;index:idx_state_value"`
	Uf           string `sql:"not null;index:idx_uf_value"`
	Logradouro   string
	Neighborhood string
	Address      string
	Complement   string
	Value        string `sql:"not null;unique_index;index:idx_cep_value"`
}
