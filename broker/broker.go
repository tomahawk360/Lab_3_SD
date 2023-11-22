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

func ensure_consistency(ctx context.Context, req *pb.GetSoldiersServiceReq, s *getSoldiersServer) int {
	index := 0

	result_1, err := s.fulcrum_client_1.AskServer(ctx, &pb.AskServerServiceReq{
		Id:     req.Id,
		Sector: req.Sector,
		Base:   "",
	})
	if err != nil {
		log.Fatal(err)
	}

	result_2, err := s.fulcrum_client_2.AskServer(ctx, &pb.AskServerServiceReq{
		Id:     req.Id,
		Sector: req.Sector,
		Base:   "",
	})
	if err != nil {
		log.Fatal(err)
	}

	result_3, err := s.fulcrum_client_3.AskServer(ctx, &pb.AskServerServiceReq{
		Id:     req.Id,
		Sector: req.Sector,
		Base:   "",
	})
	if err != nil {
		log.Fatal(err)
	}

	if result_1.Clock[1] >= result_2.Clock[1] && result_1.Clock[1] >= result_3.Clock[1] {
		index = 0
	} else if result_3.Clock[1] >= result_1.Clock[1] && result_3.Clock[1] >= result_2.Clock[1] {
		index = 2
	} else if result_2.Clock[0] >= result_1.Clock[0] && result_2.Clock[0] >= result_3.Clock[0] {
		index = 1
	} else if result_3.Clock[0] >= result_1.Clock[0] && result_3.Clock[0] >= result_2.Clock[0] {
		index = 2
	} else if result_1.Clock[2] >= result_2.Clock[2] && result_1.Clock[2] >= result_3.Clock[2] {
		index = 0
	} else if result_2.Clock[2] >= result_1.Clock[2] && result_2.Clock[2] >= result_3.Clock[2] {
		index = 1
	}

	return index
}

func (s *getSoldiersServer) GetSoldiers(ctx context.Context, req *pb.GetSoldiersServiceReq) (*pb.GetSoldiersServiceRes, error) {
	urls := [3]string{"dist113.inf.santiago.usm.cl:50051", "dist114.inf.santiago.usm.cl:50052", "dist115.inf.santiago.usm.cl:50053"}
	index := 0

	if req.Id == "Error" {
		index = ensure_consistency(ctx, req, s)
	} else {
		index = rand.Intn(2)
	}

	server_id := urls[index]
	valor := int64(0)
	clock := []int64{0, 0, 0}

	fmt.Println("Conectando a servidor: ")
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
		clock = result.Clock

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
	urls := [3]string{"dist113.inf.santiago.usm.cl:50051", "dist114.inf.santiago.usm.cl:50052", "dist115.inf.santiago.usm.cl:50053"}

	index := rand.Intn(2)
	server_id := urls[index]

	getServerRes := &pb.GetServerServiceRes{
		Id: server_id,
	}

	return getServerRes, nil
}

func main() {
	fulcrum_serv_1, err := grpc.Dial("dist113.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Fulcrum server 1: %v", err)
	}
	defer fulcrum_serv_1.Close()
	fulcrum_client_1 := pb.NewServidorServiceClient(fulcrum_serv_1)

	fulcrum_serv_2, err := grpc.Dial("dist114.inf.santiago.usm.cl:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Fulcrum server 2: %v", err)
	}
	defer fulcrum_serv_2.Close()
	fulcrum_client_2 := pb.NewServidorServiceClient(fulcrum_serv_2)

	fulcrum_serv_3, err := grpc.Dial("dist115.inf.santiago.usm.cl:50053", grpc.WithInsecure())
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
