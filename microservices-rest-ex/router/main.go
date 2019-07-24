package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/campoy/links/microservices-rest/repository"
	"github.com/campoy/links/microservices-rest/repository/client"
	"github.com/kelseyhightower/envconfig"
)

var links repository.LinkRepository

func main() {
	var config struct {
		Address    string `default:"0.0.0.0:8085"`
		Repository string `default:"0.0.0.0:8080"`
	}
	if err := envconfig.Process("ROUTER", &config); err != nil {
		log.Fatal(err)
	}

	links = client.New(config.Repository)

	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

func handleVisit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(id)
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := links.CountVisit(id); err != nil {
		log.Printf("could not track visit: %v", err)
	}

	fmt.Fprintf(w, "<p>redirecting to %s...</p>", l.URL)
	fmt.Fprintf(w, "<script>setTimeout(function() { window.location = '%s'}, 1000)</script>", l.URL)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(id)
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(l); err != nil {
		log.Printf("could not encode link information")
	}
}
