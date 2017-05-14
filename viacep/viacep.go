package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CEP type
type CEP struct {
	CEP         string `json:"cep,omitempty"`
	Logradouro  string `json:",omitempty"`
	Complemento string `json:",omitempty"`
	Bairro      string `json:",omitempty"`
	Localidade  string `json:",omitempty"`
	UF          string `json:"uf,omitempty"`
	Unidade     string `json:",omitempty"`
	IBGE        string `json:"ibge,omitempty"`
	GIA         string `json:"gia,omitempty"`
}

// Get the data of a given CEP
func Get(cep string) (result CEP, err error) {
	resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%v/json/", cep))
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
