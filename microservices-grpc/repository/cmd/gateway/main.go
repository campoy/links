package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/campoy/links/microservices-grpc/repository/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const port = ":8081"

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterRepositoryHandlerFromEndpoint(context.Background(), mux, "localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(port, mux))
}
