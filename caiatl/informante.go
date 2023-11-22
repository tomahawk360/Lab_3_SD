package main

import (
	"context"
	"fmt"
	"log"
	"main/pb"

	"google.golang.org/grpc"
)

type sectores struct {
	nombre    string
	reloj     []int64
	direccion string
}

func checkConsistency(client pb.BrokerServiceClient, registros []sectores, sector string, s *sectores) (int, string) {
	consistent := 0
	flag := 0
	new_clock := s.reloj

	str_id := "Error"

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

	return consistent, str_id
}

func getServer(client pb.BrokerServiceClient, req *pb.GetServerServiceReq) (*pb.GetServerServiceRes, error) {
	resp, err := client.GetServer(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func addBase(client pb.InformerServiceClient, req *pb.AddBaseServiceReq) (*pb.ConnectServiceRes, error) {
	resp, err := client.AddBase(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func renameBase(client pb.InformerServiceClient, req *pb.RenameBaseServiceReq) (*pb.ConnectServiceRes, error) {
	resp, err := client.RenameBase(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func updateValue(client pb.InformerServiceClient, req *pb.UpdateValueServiceReq) (*pb.ConnectServiceRes, error) {
	resp, err := client.UpdateValue(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func deleteBase(client pb.InformerServiceClient, req *pb.DeleteBaseServiceReq) (*pb.ConnectServiceRes, error) {
	resp, err := client.DeleteBase(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func main() {
	var registros []sectores

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client_brok := pb.NewBrokerServiceClient(cc)

	str_id := "Caiatl"

	for {
		fmt.Println("Ingrese comando: ")
		fmt.Println("1. Agregar Base")
		fmt.Println("2. Renombrar Base")
		fmt.Println("3. Actualizar Valor")
		fmt.Println("4. Borrar Base")
		fmt.Println("5. Salir del Programa")

		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			var sector string
			var base string
			var valor int64

			consistent := 0

			fmt.Println("Ingrese nombre del sector: ")
			fmt.Scanln(&sector)

			fmt.Println("Ingrese nombre de la base: ")
			fmt.Scanln(&base)

			fmt.Println("Ingrese valor: ")
			fmt.Scanln(&valor)

			for consistent == 0 {
				server_id, err := getServer(client_brok, &pb.GetServerServiceReq{Id: str_id})

				optz := grpc.WithInsecure()
				sj, err := grpc.Dial(server_id.Id, optz)
				if err != nil {
					log.Fatal(err)
				}
				defer cc.Close()

				client_serv := pb.NewInformerServiceClient(sj)

				result, err := addBase(client_serv, &pb.AddBaseServiceReq{
					Id:     str_id,
					Sector: sector,
					Base:   base,
					Valor:  &valor})
				if err != nil {
					log.Fatal(err)
				}

				s := &sectores{
					nombre:    sector,
					reloj:     result.Clock,
					direccion: result.Id,
				}

				consistent, str_id = checkConsistency(client_brok, registros, sector, s)
			}

		case 2:
			var sector string
			var base string
			var new_name string

			consistent := 0

			fmt.Println("Ingrese nombre del sector: ")
			fmt.Scanln(&sector)

			fmt.Println("Ingrese nombre de la base: ")
			fmt.Scanln(&base)

			fmt.Println("Ingrese el nuevo nombre: ")
			fmt.Scanln(&new_name)

			for consistent == 0 {
				server_id, err := getServer(client_brok, &pb.GetServerServiceReq{Id: str_id})

				optz := grpc.WithInsecure()
				sj, err := grpc.Dial(server_id.Id, optz)
				if err != nil {
					log.Fatal(err)
				}
				defer cc.Close()

				client_serv := pb.NewInformerServiceClient(sj)

				result, err := renameBase(client_serv, &pb.RenameBaseServiceReq{
					Id:     str_id,
					Sector: sector,
					Base:   base,
					Name:   new_name})
				if err != nil {
					log.Fatal(err)
				}

				s := &sectores{
					nombre:    sector,
					reloj:     result.Clock,
					direccion: result.Id,
				}

				consistent, str_id = checkConsistency(client_brok, registros, sector, s)
			}

		case 3:
			var sector string
			var base string
			var new_valor int64

			consistent := 0

			fmt.Println("Ingrese nombre del sector: ")
			fmt.Scanln(&sector)

			fmt.Println("Ingrese nombre de la base: ")
			fmt.Scanln(&base)

			fmt.Println("Ingrese el nuevo valor: ")
			fmt.Scanln(&new_valor)

			for consistent == 0 {
				server_id, err := getServer(client_brok, &pb.GetServerServiceReq{Id: str_id})

				optz := grpc.WithInsecure()
				sj, err := grpc.Dial(server_id.Id, optz)
				if err != nil {
					log.Fatal(err)
				}
				defer cc.Close()

				client_serv := pb.NewInformerServiceClient(sj)

				result, err := updateValue(client_serv, &pb.UpdateValueServiceReq{
					Id:     str_id,
					Sector: sector,
					Base:   base,
					Valor:  new_valor})
				if err != nil {
					log.Fatal(err)
				}

				s := &sectores{
					nombre:    sector,
					reloj:     result.Clock,
					direccion: result.Id,
				}

				consistent, str_id = checkConsistency(client_brok, registros, sector, s)
			}

		case 4:
			var sector string
			var base string

			consistent := 0

			fmt.Println("Ingrese nombre del sector: ")
			fmt.Scanln(&sector)

			fmt.Println("Ingrese nombre de la base: ")
			fmt.Scanln(&base)

			for consistent == 0 {
				server_id, err := getServer(client_brok, &pb.GetServerServiceReq{Id: str_id})

				optz := grpc.WithInsecure()
				sj, err := grpc.Dial(server_id.Id, optz)
				if err != nil {
					log.Fatal(err)
				}
				defer cc.Close()

				client_serv := pb.NewInformerServiceClient(sj)

				result, err := deleteBase(client_serv, &pb.DeleteBaseServiceReq{
					Id:     str_id,
					Sector: sector,
					Base:   base})
				if err != nil {
					log.Fatal(err)
				}

				s := &sectores{
					nombre:    sector,
					reloj:     result.Clock,
					direccion: result.Id,
				}

				consistent, str_id = checkConsistency(client_brok, registros, sector, s)
			}

		case 5:
			return // Exit the program

		default:
			fmt.Println("Ingrese una opcion valida")
			continue // Restart the loop
		}
	}
}
