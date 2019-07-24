package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	pb "github.com/campoy/links/repository/proto"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type server struct {
	links  pb.RepositoryClient
	router string
}

func main() {
	var config struct {
		Address    string `default:"0.0.0.0:8090"`
		Repository string `default:"0.0.0.0:8080"`
		Router     string `default:"loca0.0.0.05"`
	}
	if err := envconfig.Process("WEB", &config); err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	log.Printf("connecting to repository on %s", config.Repository)
	conn, err := grpc.Dial(config.Repository, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	s := server{
		links:  pb.NewRepositoryClient(conn),
		router: config.Router,
	}

	log.Printf("listening on %s", config.Address)
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
		l, err := s.links.New(r.Context(), &pb.NewRequest{Url: r.FormValue("link")})
		if err != nil {
			data.Code = http.StatusBadRequest
			data.Msg = err.Error()
		} else {
			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", s.router, l.Id)
			data.Stats = fmt.Sprintf("%s/s/%s", s.router, l.Id)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}
