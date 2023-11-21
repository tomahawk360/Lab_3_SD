package main

import (
	"context"
	"fmt"
	"log"
	"main/pb"
	"strconv"

	"google.golang.org/grpc"
)

func getSoldiers(client pb.BrokerServiceClient, req *pb.GetSoldiersServiceReq) (*pb.GetSoldiersServiceRes, error) {
	resp, err := client.GetSoldiers(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client_brok := pb.NewBrokerServiceClient(cc)

	str_id := "La Vanguardia"

	fmt.Println("Pulse cualquier tecla para conocer el numero de soldados")
	fmt.Scanln()

	var sector string
	var base string

	fmt.Println("Ingrese nombre del sector: ")
	fmt.Scanln(&sector)

	fmt.Println("Ingrese nombre de la base: ")
	fmt.Scanln(&base)

	soldiers, err := getSoldiers(client_brok, &pb.GetSoldiersServiceReq{
		Id:     str_id,
		Sector: sector,
		Base:   base,
	})

	fmt.Println(strconv.FormatInt(soldiers.Valor, 10))
}
