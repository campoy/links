package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/campoy/links/monolith/repository"
)

var (
	domain = flag.String("d", "http://localhost:8080", "URL where the server will be accessible")
	links  repository.LinkRepository
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	db, err := repository.NewDiskRepository("data")
	if err != nil {
		log.Fatal(err)
	}
	links = db

	http.HandleFunc("/", handleNew)
	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var home = template.Must(template.ParseFiles("home.html"))

func handleNew(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Code  int
		Msg   string
		Link  string
		Stats string
	}{Code: http.StatusOK}

	if r.Method == http.MethodPost {
		l, err := links.New(r.FormValue("link"))
		if err != nil {
			data.Code = http.StatusBadRequest
			data.Msg = "the given link is not a valid URL"
		} else {
			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", *domain, l.ID)
			data.Stats = fmt.Sprintf("%s/s/%s", *domain, l.ID)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
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
