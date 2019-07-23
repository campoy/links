package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/campoy/links/microservices/repository"
	"github.com/campoy/links/microservices/repository/client"
)

var links = client.New("http://localhost:8080")

func main() {
	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(":8085", nil))
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

	links.CountVisit(id)

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
