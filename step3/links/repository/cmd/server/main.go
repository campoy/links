package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/campoy/links/repository"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

var links repository.LinkRepository

func main() {
	var config struct {
		Address string `default:"0.0.0.0:8080"`
	}
	if err := envconfig.Process("REPOSITORY", &config); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())

	db, err := repository.NewDiskRepository("data")
	if err != nil {
		log.Fatal(err)
	}
	links = db

	mux := mux.NewRouter()
	mux.HandleFunc("/link/", newLink).Methods("POST")
	mux.HandleFunc("/link/{id}", getLink).Methods("GET")
	mux.HandleFunc("/link/{id}", countVisit).Methods("POST")
	log.Fatal(http.ListenAndServe(config.Address, mux))
}

func newLink(w http.ResponseWriter, r *http.Request) {
	var data struct{ URL string }
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := links.New(data.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(link); err != nil {
		log.Printf("could not encode link: %v", err)
	}
}

func getLink(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	link, err := links.Get(id)
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(link); err != nil {
		log.Printf("could not encode link: %v", err)
	}
}

func countVisit(w http.ResponseWriter, r *http.Request) {
	// TODO(you): Implement the body of this function.
	log.Fatal("NOT IMPLEMENTED")
}
