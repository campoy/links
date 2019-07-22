package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", handleNew)
	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type link struct {
	id    string
	url   string
	count int
}

var links = make(map[string]link)

var home = template.Must(template.ParseFiles("home.html"))

func handleNew(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Method)
	var data struct {
		Code int
		Msg  string
		ID   string
	}

	if r.Method == http.MethodPost {
		u := r.FormValue("l")
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
			links[id] = link{id: id, url: u}

			data.Code = http.StatusCreated
			data.Msg = "link successfully created!"
			data.ID = id
		}
	}
	log.Printf("rendering data: %v", data)
	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

func handleVisit(w http.ResponseWriter, r *http.Request) {
	l, ok := links[r.URL.Path[3:]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	l.count++
	links[l.id] = l

	http.Redirect(w, r, l.url, http.StatusPermanentRedirect)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	l, ok := links[r.URL.Path[3:]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "link visited %d times\n", l.count)
}
