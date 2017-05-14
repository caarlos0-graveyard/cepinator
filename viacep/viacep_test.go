package viacep_test

import (
	"testing"

	"github.com/caarlos0/cepinator/viacep"
	"github.com/stretchr/testify/assert"
)

func TestViacep(t *testing.T) {
	var assert = assert.New(t)
	result, err := viacep.Get("01001000")
	assert.NoError(err)
	assert.Equal(viacep.CEP{
		CEP:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		UF:          "SP",
		IBGE:        "3550308",
		GIA:         "1004",
	}, result)
}
