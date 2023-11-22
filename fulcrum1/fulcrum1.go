package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
    "strings"
	"context"
	"log"
	"main/pb"
	"strconv"
	"net"

	"google.golang.org/grpc"
)

type Action struct {
    Cmd        string
	Sector	   string      
    Base       string
    NewValue   string
}

type server struct {
	pb.LogServiceServer
}

type InformerServiceServer struct{
	pb.UnimplementedInformerServiceServer
}

type ServidorServiceServer struct{
	pb.UnimplementedServidorServiceServer
}

var (
    logfile, _ = os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    logger     = log.New(logfile, "", 0)
)

func requestLogs() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) 
	if err != nil {
		fmt.Println("error")
	}
	defer conn.Close()
	c := pb.NewLogServiceClient(conn)
	r, err := c.SendLogs(context.Background(), &pb.LogRequest{ServerID: ""})
	if err != nil {
		log.Fatalf("Could not get logs: %v", err)
	}


	conn2, err := grpc.Dial("localhost:50050", grpc.WithInsecure()) 
	if err != nil {
		fmt.Println("error")
	}
	defer conn2.Close()
	c2 := pb.NewLogServiceClient(conn2)
	r2, err := c2.SendLogs(context.Background(), &pb.LogRequest{ServerID: ""})
	if err != nil {
		log.Fatalf("Could not get logs: %v", err)
	}
	processLogs(r.Logs, r2.Logs)
}

func processLogs(logs2 []string, logs3 []string) {
	file, err := os.Open("logfile.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

}

func GetSoldiers (file *os.File, base string) (int64){
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, base) {
			//split the line
			//return the value

			stringSlice := strings.Fields(line)
			i, err := strconv.ParseInt(stringSlice[1], 10, 64)
			if err != nil {
				panic(err)
			}
			return i
		}
	}
	return -1
}


func (s *ServidorServiceServer) AskServer (ctx context.Context, msg *pb.AskServerServiceReq) (*pb.AskServerServiceRes, error){
	sectorFile, err := os.OpenFile(msg.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

	var soldier int64 = GetSoldiers(sectorFile, msg.Base)

	return &pb.AskServerServiceRes{Id: "1", Valor: soldier, Clock: []int64{0,0,0}}, nil
}

func (s *InformerServiceServer) DeleteBase (ctx context.Context, msg *pb.DeleteBaseServiceReq) (*pb.ConnectServiceRes, error){
	sectorFile, err := os.OpenFile(msg.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

	removeBase(sectorFile, msg.Base)
    fmt.Printf("Quitando base \"%v\" en sector \"%v\"\n", msg.Base, msg.Sector)
	logger.Printf("%s BorrarBase %v %v", time.Now().Format("15:04:05"), msg.Sector, msg.Base)

	return &pb.ConnectServiceRes{Id: "1", Clock: []int64{0,0,0}}, nil
}

func (s *InformerServiceServer) UpdateValue (ctx context.Context, msg *pb.UpdateValueServiceReq) (*pb.ConnectServiceRes, error) {
	sectorFile, err := os.OpenFile(msg.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

    writer := bufio.NewWriter(sectorFile)

	if updateBaseValue(sectorFile, msg.Base, msg.Valor) {
		fmt.Printf("Actualizando base \"%v\" en sector \"%v\" a valor %v\n", msg.Base, msg.Sector, msg.Valor)
		logger.Printf("%s ActualizarValor %v %v %v", time.Now().Format("15:04:05"), msg.Sector, msg.Base, msg.Valor)
	} else {
		fmt.Printf("Añadiendo base \"%v\" en sector \"%v\" con valor %v\n", msg.Base, msg.Sector, msg.Valor)
		logger.Printf("%s AgregarBase %v %v %v", time.Now().Format("15:04:05"), msg.Sector, msg.Base, msg.Valor)
		fmt.Fprintf(writer, "%v %v\n", msg.Base, msg.Valor)
		writer.Flush()
	}
	return &pb.ConnectServiceRes{Id: "1", Clock: []int64{0,0,0}}, nil
}

func (s *InformerServiceServer) RenameBase (ctx context.Context, msg *pb.RenameBaseServiceReq) (*pb.ConnectServiceRes, error) {
	sectorFile, err := os.OpenFile(msg.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

	renameBases(sectorFile, msg.Base, msg.Name)
	fmt.Printf("Renombrando base \"%v\" en sector \"%v\" a valor %v\n", msg.Base , msg.Sector, msg.Name)
	logger.Printf("%s RenombrarBase %v %v %v", time.Now().Format("15:04:05"), msg.Sector, msg.Base, msg.Name)

	return &pb.ConnectServiceRes{Id: "1", Clock: []int64{0,0,0}}, nil
}

func (s *InformerServiceServer) AddBase (ctx context.Context, msg *pb.AddBaseServiceReq) (*pb.ConnectServiceRes, error) {
    sectorFile, err := os.OpenFile(msg.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

    writer := bufio.NewWriter(sectorFile)

	if !baseExists(sectorFile, msg.Base) {
		fmt.Printf("Añadiendo base \"%v\" en sector \"%v\" con valor %v\n", msg.Base, msg.Sector, msg.Valor)
		logger.Printf("%s AgregarBase %v %v %v", time.Now().Format("15:04:05"), msg.Sector, msg.Base, msg.Valor)
		fmt.Fprintf(writer, "%v %v\n", msg.Base, msg.Valor)
	} else {
		fmt.Printf("La base \"%v\" ya existe en el sector \"%v\"\n", msg.Base, msg.Sector)
	}
	writer.Flush()
	return &pb.ConnectServiceRes{Id: "1", Clock: []int64{0,0,0}}, nil
}

func (s *server) SendLogs(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
    logs, err := readLogs() 
    if err != nil {
        return nil, err 
    }
    return &pb.LogResponse{Logs: logs}, nil
}

func readLogs() ([]string, error) {
	file, err := os.Open("logfile.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func baseExists(file *os.File, base string) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, base) {
			return true
		}
	}
	return false
}

func updateBaseValue(file *os.File, base string, newValue int64) bool {
	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("Error reading lines: %v", err)
		return false
	}

	for i, line := range lines {
		if strings.Contains(line,base) {
			lines[i] = fmt.Sprintf("%v %v",base, newValue)
			writeLines(file, lines)
			return true
		}
	}

	return false
}

func renameBases(file *os.File, oldBase string, newBase string) {
	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("Error reading lines: %v", err)
		return
	}

	baseExists := false

	var oldValue []string
	var index int

	for i, line := range lines {
		if strings.Contains(line, oldBase) {
			oldValue = strings.Fields(line)
			index = i
			baseExists = true
			break
		}
	}

	if baseExists {
		fmt.Println("existe")
		lines[index] = fmt.Sprintf("%s %s", newBase, oldValue[1])
	} else {
		fmt.Println("no existe")
		lines = append(lines, fmt.Sprintf("%s 0", newBase))
	}
	writeLines(file, lines)
}


func removeBase(file *os.File, base string) {
	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("Error reading lines: %v", err)
		return
	}

	var updatedLines []string
	for _, line := range lines {
		if !strings.Contains(line, base) {
			updatedLines = append(updatedLines, line)
		}
	}

	writeLines(file, updatedLines)
}



func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func writeLines(file *os.File, lines []string) {
	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}


func main() {
    defer logfile.Close()
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterInformerServiceServer(grpcServer, &InformerServiceServer{})
	log.Printf("Server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
    }

	//implemnet listen for ask

	list, err := net.Listen("tcp", ":50040")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer2 := grpc.NewServer()
	pb.RegisterServidorServiceServer(grpcServer2, &ServidorServiceServer{})
	log.Printf("Server listening on %v", list.Addr())

	if err := grpcServer2.Serve(list); err != nil {
			log.Fatalf("Failed to serve: %v", err)
	}
	
}
