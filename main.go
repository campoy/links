package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var domain = flag.String("d", "http://localhost:8080", "URL where the server will be accessible")

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", handleNew)
	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type link struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

var links = make(map[string]link)

func handleNew(w http.ResponseWriter, r *http.Request) {
	var home = template.Must(template.ParseFiles("home.html"))

	data := struct {
		Code  int
		Msg   string
		Link  string
		Stats string
	}{Code: http.StatusOK}

	if r.Method == http.MethodPost {
		u := r.FormValue("link")

		if _, err := url.ParseRequestURI(u); err != nil {
			data.Code = http.StatusBadRequest
			data.Msg = "the given link is not a valid URL"
		} else {
			id := randomString()
			for {
				if _, ok := links[id]; !ok {
					break
				}
				id = randomString()
			}
			links[id] = link{ID: id, URL: u}

			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", *domain, id)
			data.Stats = fmt.Sprintf("%s/s/%s", *domain, id)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

func handleVisit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, ok := links[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	l.Count++
	links[l.ID] = l

	fmt.Fprintf(w, "<p>redirecting to %s...</p>", l.URL)
	fmt.Fprintf(w, "<script>setTimeout(function() { window.location = '%s'}, 1000)</script>", l.URL)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, ok := links[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(l); err != nil {
		log.Printf("could not encode link information")
	}
}
