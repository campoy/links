package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/campoy/links/microservices-grpc/repository"
	pb "github.com/campoy/links/microservices-grpc/repository/proto"
	"google.golang.org/grpc"
)

var links pb.RepositoryClient

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	links = pb.NewRepositoryClient(conn)

	http.HandleFunc("/l/", handleVisit)
	http.HandleFunc("/s/", handleStats)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func handleVisit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(r.Context(), &pb.IDRequest{Id: id})
	if err != nil {
		if err == repository.ErrNoSuchLink {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_, err = links.CountVisit(r.Context(), &pb.IDRequest{Id: id})
	if err != nil {
		log.Printf("could not count visit: %v", err)
	}

	fmt.Fprintf(w, "<p>redirecting to %s...</p>", l.Url)
	fmt.Fprintf(w, "<script>setTimeout(function() { window.location = '%s'}, 1000)</script>", l.Url)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[3:]
	l, err := links.Get(r.Context(), &pb.IDRequest{Id: id})
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