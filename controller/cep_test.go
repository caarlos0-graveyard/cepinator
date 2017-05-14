package controller_test

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"encoding/json"

	"github.com/caarlos0/cepinator/cache"
	"github.com/caarlos0/cepinator/controller"
	"github.com/caarlos0/cepinator/viacep"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	var assert = assert.New(t)
	var cache = cache.New(":6379")
	defer cache.Close()
	var r = mux.NewRouter()
	r.HandleFunc("/{cep}", controller.CEP(cache))
	var srv = httptest.NewServer(r)
	defer srv.Close()

	var result viacep.CEP
	resp, err := http.Get(fmt.Sprintf("%v/%v", srv.URL, "01001-000"))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal("01001-000", result.CEP)
	assert.Equal("Praça da Sé", result.Logradouro)
}

func TestControllerNoRedis(t *testing.T) {
	var assert = assert.New(t)
	var cache = cache.New(":meh")
	defer cache.Close()
	var r = mux.NewRouter()
	r.HandleFunc("/{cep}", controller.CEP(cache))
	var srv = httptest.NewServer(r)
	defer srv.Close()

	var result viacep.CEP
	resp, err := http.Get(fmt.Sprintf("%v/%v", srv.URL, "01001-000"))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal("01001-000", result.CEP)
	assert.Equal("Praça da Sé", result.Logradouro)
}
