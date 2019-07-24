package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/campoy/links/microservices-rest/repository"
	"github.com/campoy/links/microservices-rest/repository/client"
	"github.com/kelseyhightower/envconfig"
)

type server struct {
	links  repository.LinkRepository
	router string
}

func main() {
	var config struct {
		Address    string `default:"localhost:8090"`
		Repository string `default:"http://localhost:8080"`
		Router     string `default:"http://localhost:8085"`
	}
	if err := envconfig.Process("WEB", &config); err != nil {
		log.Fatal(err)
	}

	s := server{links: client.New(config.Repository),
		router: config.Router,
	}

	http.HandleFunc("/", s.handleNew)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

var home = template.Must(template.ParseFiles("home.html"))

func (s *server) handleNew(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Code  int
		Msg   string
		Link  string
		Stats string
	}{Code: http.StatusOK}

	if r.Method == http.MethodPost {
		l, err := s.links.New(r.FormValue("link"))
		if err != nil {
			data.Code = http.StatusBadRequest
			data.Msg = err.Error()
		} else {
			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", s.router, l.ID)
			data.Stats = fmt.Sprintf("%s/s/%s", s.router, l.ID)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}
