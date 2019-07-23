package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	pb "github.com/campoy/links/microservices-grpc/repository/proto"
	"google.golang.org/grpc"
)

var (
	router = flag.String("d", "http://localhost:8085", "URL where the router will be accessible")
	links  pb.RepositoryClient
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	links = pb.NewRepositoryClient(conn)

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
		l, err := links.New(r.Context(), &pb.NewRequest{Url: r.FormValue("link")})
		if err != nil {
			data.Code = http.StatusBadRequest
			data.Msg = err.Error()
		} else {
			data.Code = http.StatusCreated
			data.Link = fmt.Sprintf("%s/l/%s", *router, l.Id)
			data.Stats = fmt.Sprintf("%s/s/%s", *router, l.Id)
		}
	}

	if err := home.Execute(w, data); err != nil {
		log.Printf("could not render template: %v", err)
	}
}
