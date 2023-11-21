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

type getSoldiersServer struct {
	pb.BrokerServiceServer
	fulcrum_client_1 pb.ServidorServiceClient
	fulcrum_client_2 pb.ServidorServiceClient
	fulcrum_client_3 pb.ServidorServiceClient
}

func (s *getSoldiersServer) GetSoldiers(ctx context.Context, req *pb.GetSoldiersServiceReq) (*pb.GetSoldiersServiceRes, error) {
	urls := [3]string{"localhost:8010", "localhost:8012", "localhost:8014"}

	index := rand.Intn(2)
	server_id := urls[index]

	valor := int64(0)
	clock := []int64{0, 0, 0}

	fmt.Println(server_id)

	switch index {
	case 0:
		result, err := s.fulcrum_client_1.AskServer(ctx, &pb.AskServerServiceReq{
			Id:     req.Id,
			Sector: req.Sector,
			Base:   req.Base,
		})
		if err != nil {
			log.Fatal(err)
		}

		valor = result.Valor
		clock[0] = int64(1)

	case 1:
		result, err := s.fulcrum_client_2.AskServer(ctx, &pb.AskServerServiceReq{
			Id:     req.Id,
			Sector: req.Sector,
			Base:   req.Base,
		})
		if err != nil {
			log.Fatal(err)
		}

		valor = result.Valor
		clock[1] = int64(1)

	case 2:
		result, err := s.fulcrum_client_3.AskServer(ctx, &pb.AskServerServiceReq{
			Id:     req.Id,
			Sector: req.Sector,
			Base:   req.Base,
		})
		if err != nil {
			log.Fatal(err)
		}

		valor = result.Valor
		clock[2] = int64(1)

	default:
		log.Fatal("Index no valido")
	}

	getSoldiersRes := &pb.GetSoldiersServiceRes{
		Id:    server_id,
		Valor: valor,
		Clock: clock,
	}

	return getSoldiersRes, nil
}

func (s *getSoldiersServer) GetServer(ctx context.Context, req *pb.GetServerServiceReq) (*pb.GetServerServiceRes, error) {
	urls := [3]string{"localhost:8010", "localhost:8012", "localhost:8014"}

	index := rand.Intn(2)
	server_id := urls[index]

	getServerRes := &pb.GetServerServiceRes{
		Id: server_id,
	}

	return getServerRes, nil
}

func main() {
	fulcrum_serv_1, err := grpc.Dial("localhost:8010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Fulcrum server 1: %v", err)
	}
	defer fulcrum_serv_1.Close()
	fulcrum_client_1 := pb.NewServidorServiceClient(fulcrum_serv_1)

	fulcrum_serv_2, err := grpc.Dial("localhost:8012", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Fulcrum server 2: %v", err)
	}
	defer fulcrum_serv_2.Close()
	fulcrum_client_2 := pb.NewServidorServiceClient(fulcrum_serv_2)

	fulcrum_serv_3, err := grpc.Dial("localhost:8014", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Fulcrum server 3: %v", err)
	}
	defer fulcrum_serv_3.Close()
	fulcrum_client_3 := pb.NewServidorServiceClient(fulcrum_serv_3)

	getSoldiersServer := &getSoldiersServer{
		fulcrum_client_1: fulcrum_client_1,
		fulcrum_client_2: fulcrum_client_2,
		fulcrum_client_3: fulcrum_client_3,
	}

	getServerListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port: 8080")

	s := grpc.NewServer()
	pb.RegisterBrokerServiceServer(s, getSoldiersServer)

	if err := s.Serve(getServerListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
