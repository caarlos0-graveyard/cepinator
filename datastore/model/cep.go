package model

type CEP struct {
	ID           int64  `db:"id"`
	City         string `db:"city"`
	State        string `db:"state"`
	Uf           string `db:"uf"`
	Logradouro   string `db:"logradouro"`
	Neighborhood string `db:"neighborhood"`
	Address      string `db:"address"`
	Complement   string `db:"complement"`
	Value        string `db:"value"`
}
