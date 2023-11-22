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

	cc, err := grpc.Dial("dist116.inf.santiago.usm.cl:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client_brok := pb.NewBrokerServiceClient(cc)

	str_id := "La Vanguardia"
	consistent := 0

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

			for consistent == 0 {
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
				str_id = "Error"
				new_clock := s.reloj

				for i, v := range registros {
					if v.nombre == sector {
						saved_clock := v.reloj

						if new_clock[0] >= saved_clock[0] && new_clock[1] >= saved_clock[1] && new_clock[2] >= saved_clock[2] {
							registros[i] = *s
							flag = 1
							consistent = 1
						}

					}
				}

				if flag == 0 {
					registros = append(registros, *s)
					consistent = 1
				}
			}
		} else {
			return
		}
	}
}
