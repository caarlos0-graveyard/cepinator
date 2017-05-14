package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/apex/log"
	"github.com/caarlos0/cepinator/cache"
	"github.com/caarlos0/cepinator/viacep"
	"github.com/gorilla/mux"
)

// CEP controller
func CEP(cache cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cep = string(regexp.MustCompile("[^0-9]").ReplaceAll(
			[]byte(mux.Vars(r)["cep"]),
			[]byte{},
		))
		var key = fmt.Sprintf("cepinator:%v", cep)
		var log = log.WithField("cep", cep)
		var result viacep.CEP
		if err := cache.Get(key, &result); err == nil {
			log.Info("found in cache")
			if err := json.NewEncoder(w).Encode(result); err != nil {
				log.WithError(err).Error("failed to encode cached result")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		result, err := viacep.Get(cep)
		if err != nil {
			log.WithError(err).Error("failed to get from viacep")
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		if err := cache.Put(key, result); err != nil {
			log.WithError(err).Error("failed to cache viacep result")
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.WithError(err).Error("failed to encode viacep result")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
