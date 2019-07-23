package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/campoy/links/microservices/repository/client"
)

var (
	router = flag.String("d", "http://localhost:8085", "URL where the router will be accessible")
	links  = client.New("http://localhost:8080")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", handleNew)
	log.Fatal(http.ListenAndServe(":8090", nil))
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
			data.Msg = err.Error()
		} else {
			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", *router, l.ID)
			data.Stats = fmt.Sprintf("%s/s/%s", *router, l.ID)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}
