package main

import (
	"context"
	"fmt"
	"log"
	"main/pb"
	"strconv"

	"google.golang.org/grpc"
)

type sectores struct {
	nombre    string
	reloj     []int64
	direccion string
}

func getSoldiers(client pb.BrokerServiceClient, req *pb.GetSoldiersServiceReq) (*pb.GetSoldiersServiceRes, error) {
	resp, err := client.GetSoldiers(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func main() {
	var registros []sectores

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client_brok := pb.NewBrokerServiceClient(cc)

	str_id := "La Vanguardia"

	for {
		var opt string
		var sector string
		var base string

		fmt.Println("Pulse x para salir")
		fmt.Println("Pulse cualquier tecla otra tecla para conocer el numero de soldados")
		fmt.Scanln(&opt)

		if opt != "x" {

			fmt.Println("Ingrese nombre del sector: ")
			fmt.Scanln(&sector)

			fmt.Println("Ingrese nombre de la base: ")
			fmt.Scanln(&base)

			soldiers, err := getSoldiers(client_brok, &pb.GetSoldiersServiceReq{
				Id:     str_id,
				Sector: sector,
				Base:   base,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(strconv.FormatInt(soldiers.Valor, 10))

			s := &sectores{
				nombre:    sector,
				reloj:     soldiers.Clock,
				direccion: soldiers.Id,
			}

			flag := 0
			for i, v := range registros {
				if v.nombre == sector {
					registros[i] = *s
					flag = 1
				}
			}

			if flag == 0 {
				registros = append(registros, *s)
			}
		} else {
			return
		}
	}
}
