package model

import "time"

type CEP struct {
	Value        string    `db:"value" json:"value,omitempty"`
	City         string    `db:"city" json:"city,omitempty"`
	State        string    `db:"state" json:"state,omitempty"`
	Uf           string    `db:"uf" json:"uf,omitempty"`
	Logradouro   string    `db:"logradouro" json:"logradouro,omitempty"`
	Neighborhood string    `db:"neighborhood" json:"neighborhood,omitempty"`
	Address      string    `db:"address" json:"address,omitempty"`
	Complement   string    `db:"complement" json:"complement,omitempty"`
	UpdateAt     time.Time `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt    time.Time `db:"created_at" json:"created_at,omitempty"`
}
