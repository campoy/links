package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/campoy/links/repository"
	pb "github.com/campoy/links/repository/proto"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type server struct {
	links repository.LinkRepository
}

func main() {
	var config struct {
		Address string `default:"0.0.0.0:8080"`
	}
	if err := envconfig.Process("REPOSITORY", &config); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := repository.NewDiskRepository("data")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on %s", config.Address)
	s := grpc.NewServer()
	pb.RegisterRepositoryServer(s, &server{db})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *server) New(ctx context.Context, req *pb.NewRequest) (*pb.Link, error) {
	// TODO(you): implement
	return nil, errors.New("not implemented")
}

func (s *server) Get(ctx context.Context, req *pb.IDRequest) (*pb.Link, error) {
	// TODO(you): implement
	return nil, errors.New("not implemented")
}

func (s *server) CountVisit(ctx context.Context, req *pb.IDRequest) (*pb.Nothing, error) {
	// TODO(you): implement
	return nil, errors.New("not implemented")
}
