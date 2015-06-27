package main

import (
	"github.com/jinzhu/gorm"
)

type City struct {
	gorm.Model
	State State  `sql:"not null"`
	Name  string `sql:"not null;unique"`
}
