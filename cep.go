package main

import (
	"github.com/jinzhu/gorm"
)

type Cep struct {
	gorm.Model
	City         City `sql:"not null"`
	Logradouro   string
	Neighborhood string
	Address      string
	Complement   string
	Value        string `sql:"not null;unique"`
}
