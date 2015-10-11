package main

import (
	"fmt"
	"net/http"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
)

func init() {
	log.SetHandler(logfmt.Default)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(ok))
	mux.Handle("/error", http.HandlerFunc(err))

	log.Info("starting")
	err := http.ListenAndServe(":5000", httplog.New(mux))
	log.WithError(err).Fatal("error listening")
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func err(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad Request")
}
