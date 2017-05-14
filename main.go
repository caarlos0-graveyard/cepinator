package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/caarlos0/cepinator/cache"
	"github.com/caarlos0/cepinator/controller"
	"github.com/caarlos0/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var version = "dev"

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
	var log = log.WithField("version", version).
		WithField("port", config.Port).
		WithField("redis", config.RedisURL)

	var cache = cache.New(config.RedisURL)
	defer func() {
		if err := cache.Close(); err != nil {
			log.WithError(err).Error("failed to close cache")
		}
	}()

	var r = mux.NewRouter()

	r.HandleFunc("/{cep}", controller.CEP(cache))

	var srv = &http.Server{
		Handler:      httplog.New(handlers.CompressHandler(r)),
		Addr:         fmt.Sprintf("0.0.0.0:%v", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Info("starting up...")
	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).Error("failed to start up server")
	}
}
