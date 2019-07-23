package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/campoy/links/microservices-grpc/repository"
	pb "github.com/campoy/links/microservices-grpc/repository/proto"
	"google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	links repository.LinkRepository
}

func main() {
	rand.Seed(time.Now().Unix())

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := repository.NewDiskRepository("data")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterRepositoryServer(s, &server{db})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *server) New(ctx context.Context, req *pb.NewRequest) (*pb.Link, error) {
	link, err := s.links.New(req.GetUrl())
	if err != nil {
		return nil, err
	}
	return &pb.Link{
		Id:    link.ID,
		Url:   link.URL,
		Count: int64(link.Count),
	}, nil
}

func (s *server) Get(ctx context.Context, req *pb.IDRequest) (*pb.Link, error) {
	link, err := s.links.Get(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.Link{
		Id:    link.ID,
		Url:   link.URL,
		Count: int64(link.Count),
	}, nil
}

func (s *server) CountVisit(ctx context.Context, req *pb.IDRequest) (*pb.Nothing, error) {
	return nil, s.links.CountVisit(req.GetId())
}
