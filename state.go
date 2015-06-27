package main

import (
	"github.com/jinzhu/gorm"
)

type State struct {
	gorm.Model
	Abbreviation string `sql:"not null;unique"`
	Name         string `sql:"not null;unique"`
}
