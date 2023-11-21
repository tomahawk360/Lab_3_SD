package main

import (
	"fmt"
    "time"
	"context"
	"log"
	"main/pb"

	"google.golang.org/grpc"
)

type Action struct {
    Cmd        string
	Sector	   string      
    Base       string
    NewValue   string
}

type InformerServiceServer struct{
	pb.UnimplementedInformerServiceServer
}

func main(){
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewInformerServiceClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	
	request := &pb.DeleteBaseServiceReq{
		Id: "1",
		Sector: "your_sector",
		Base:   "your_base",
	}
	
	response, err := client.DeleteBase(ctx, request)
	if err != nil {
		log.Fatalf("Error calling AgregarBase: %v", err)
	}
	fmt.Printf("Response from server: %v\n", response)
	
}