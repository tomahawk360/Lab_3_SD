package main

import (
	"context"
	"fmt"
	"log"
	"main/pb"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.BrokerServiceServer
}

func (s *server) GetSoldiers(ctx context.Context, req *pb.GetSoldiersServiceReq) (*pb.GetSoldiersServiceRes, error) {
	urls := [3]string{"localhost:8010", "localhost:8012", "localhost:8014"}

	index := rand.Intn(2)
	server_id := urls[index]

	getServerRes := &pb.GetSoldiersServiceRes{
		Id: server_id,
	}

	return getServerRes, nil
}

func (s *server) GetServer(ctx context.Context, req *pb.GetServerServiceReq) (*pb.GetServerServiceRes, error) {
	urls := [3]string{"localhost:8010", "localhost:8012", "localhost:8014"}

	index := rand.Intn(2)
	server_id := urls[index]

	getServerRes := &pb.GetServerServiceRes{
		Id: server_id,
	}

	return getServerRes, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port :8080")

	s := grpc.NewServer()

	pb.RegisterBrokerServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
