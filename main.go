package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/caarlos0/cepinator/cache"
	"github.com/caarlos0/cepinator/viacep"
	"github.com/caarlos0/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Config type
type Config struct {
	Port     string `env:"PORT" envDefault:"3000"`
	RedisURL string `env:"REDIS_URL" envDefault:":6379"`
}

func init() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)
}

func main() {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}

	var cache = cache.New(config.RedisURL)
	defer cache.Close()

	var r = mux.NewRouter()

	r.HandleFunc("/{cep}", func(w http.ResponseWriter, r *http.Request) {
		var cep = strings.Replace(mux.Vars(r)["cep"], "[^0-9]", "", -1)
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
	})

	var srv = &http.Server{
		Handler:      httplog.New(handlers.CompressHandler(r)),
		Addr:         fmt.Sprintf("0.0.0.0:%v", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.WithField("port", config.Port).Info("starting up...")
	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).Error("failed to start up server")
	}
}
