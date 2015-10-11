package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	"strings"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/caarlos0/env"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Config type
type Config struct {
	Port     string `env:"PORT" envDefault:"3000"`
	RedisURL string `env:"REDIS_URL" envDefault:":6379"`
}

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

var config Config

func init() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)
	if err := env.Parse(&config); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
}

func main() {
	ring, codec := setupCache()
	defer ring.Close()

	var r = mux.NewRouter()

	r.HandleFunc("/{cep}", func(w http.ResponseWriter, r *http.Request) {
		var cep = strings.Replace(mux.Vars(r)["cep"], "[^0-9]", "", -1)
		var log = log.WithField("cep", cep)
		var result CEP
		if err := codec.Get(cep, &result); err == nil {
			log.Info("found in cache")
			if err := json.NewEncoder(w).Encode(result); err != nil {
				log.WithError(err).Error("failed to encode cached result")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%v/json/", cep))
		if err != nil {
			log.WithError(err).Error("failed to get from viacep")
			http.Error(w, err.Error(), resp.StatusCode)
			return
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.WithError(err).Error("failed to decode viacep result")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := codec.Set(&cache.Item{
			Key:        cep,
			Object:     result,
			Expiration: time.Hour * 24 * 30, // 1mo
		}); err != nil {
			log.WithError(err).Error("failed to cache viacep result")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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

func setupCache() (*redis.Ring, *cache.Codec) {
	var ring = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": config.RedisURL,
		},
	})
	var codec = &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	return ring, codec
}
