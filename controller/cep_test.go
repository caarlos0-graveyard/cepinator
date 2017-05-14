package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caarlos0/cepinator/cache"
	"github.com/caarlos0/cepinator/controller"
	"github.com/caarlos0/cepinator/viacep"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	for _, url := range []string{":6379", "foo.none:6379"} {
		t.Run(url, func(t *testing.T) {
			var assert = assert.New(t)
			var cache = cache.New(":6379")
			defer func() {
				_ = cache.Close()
			}()
			var r = mux.NewRouter()
			r.HandleFunc("/{cep}", controller.CEP(cache))
			var srv = httptest.NewServer(r)
			defer srv.Close()

			// delete from cache first
			_ = cache.Delete("cepinator:01001000")

			// several requests
			for i := 0; i < 10; i++ {
				var result viacep.CEP
				resp, err := http.Get(fmt.Sprintf("%v/%v", srv.URL, "01001-000"))
				assert.NoError(err)
				assert.Equal(http.StatusOK, resp.StatusCode)
				_ = json.NewDecoder(resp.Body).Decode(&result)

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
		})
	}
}
