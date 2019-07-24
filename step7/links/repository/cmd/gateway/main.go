package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/campoy/links/repository/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

func main() {
	var config struct {
		Address    string `default:"0.0.0.0:8081"`
		Repository string `default:"0.0.0.0:8080"`
	}
	if err := envconfig.Process("GATEWAY", &config); err != nil {
		log.Fatal(err)
	}

	mux := runtime.NewServeMux()

	// TODO(you): create a gRPC gateway and register it on mux
	// info: https://github.com/grpc-ecosystem/grpc-gateway

	log.Printf("serving gRPC endpoint from %s on REST %s", config.Repository, config.Address)
	log.Fatal(http.ListenAndServe(config.Address, mux))
}
