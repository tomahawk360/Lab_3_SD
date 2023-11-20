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
	"net"

	"google.golang.org/grpc"
)

type Action struct {
    Cmd        string
	Sector	   string      
    Base       string
    NewValue   string
}

var (
    logfile, _ = os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    logger     = log.New(logfile, "", 0)
)

type server struct {
	pb.LogServiceServer
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

func handleAction(action Action) {
	sectorFile, err := os.OpenFile(action.Sector+".txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening sector file: %v", err)
    }
    defer sectorFile.Close()

    writer := bufio.NewWriter(sectorFile)

    switch action.Cmd {
    case "AgregarBase":
		if !baseExists(sectorFile, action.Base) {
			fmt.Printf("Añadiendo base \"%v\" en sector \"%v\" con valor %v\n", action.Base, action.Sector, action.NewValue)
			logger.Printf("%s AgregarBase %v %v %v", time.Now().Format("15:04:05"), action.Sector, action.Base, action.NewValue)
			fmt.Fprintf(writer, "%v %v\n", action.Base, action.NewValue)
		} else {
			fmt.Printf("La base \"%v\" ya existe en el sector \"%v\"\n", action.Base, action.Sector)
		}

    case "ActualizarValor":
		if updateBaseValue(sectorFile, action.Base, action.NewValue) {
			fmt.Printf("Actualizando base \"%v\" en sector \"%v\" a valor %v\n", action.Base, action.Sector, action.NewValue)
			logger.Printf("%s ActualizarValor %v %v %v", time.Now().Format("15:04:05"), action.Sector, action.Base, action.NewValue)
		} else {
			fmt.Printf("Añadiendo base \"%v\" en sector \"%v\" con valor %v\n", action.Base, action.Sector, action.NewValue)
			logger.Printf("%s AgregarBase %v %v %v", time.Now().Format("15:04:05"), action.Sector, action.Base, action.NewValue)
			fmt.Fprintf(writer, "%v %v\n", action.Base, action.NewValue)
			writer.Flush()
		}

    case "RenombrarBase":
		renameBase(sectorFile, action.Base, action.NewValue)
        fmt.Printf("Renombrando base \"%v\" en sector \"%v\" a valor %v\n", action.Base , action.Sector, action.NewValue)
		logger.Printf("%s RenombrarBase %v %v %v", time.Now().Format("15:04:05"), action.Sector, action.Base, action.NewValue)

    case "BorrarBase":
		removeBase(sectorFile, action.Base)
        fmt.Printf("Quitando base \"%v\" en sector \"%v\"\n", action.Base, action.Sector)
		logger.Printf("%s BorrarBase %v %v", time.Now().Format("15:04:05"), action.Sector, action.Base)

	case "sync":
		break

    default:
        fmt.Println("Accion Invalida")
    }
	writer.Flush()
}

func baseExists(file *os.File, base string) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == base {
			return true
		}
	}
	return false
}

func updateBaseValue(file *os.File, base string, newValue string) bool {
	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("Error reading lines: %v", err)
		return false
	}

	for i, line := range lines {
		if strings.Contains(line,base) {
			lines[i] = fmt.Sprintf("%v %s",base, newValue)
			writeLines(file, lines)
			return true
		}
	}

	return false
}

func renameBase(file *os.File, oldBase string, newBase string) {
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
	fmt.Println("Huh")
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port :50050")

	s := grpc.NewServer()
	pb.RegisterLogServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

    defer logfile.Close()
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        text, _ := reader.ReadString('\n')
        text = strings.TrimSuffix(text, "\n")
        commands := strings.Split(text, " ")
        if len(commands) < 3 {
            fmt.Println("Comando invalido, debe tener formato <Accion> <Sector> <Base> <Valor>")
            continue
        }

        action := Action{Cmd: commands[0], Sector: commands[1], Base: commands[2]}

        if len(commands) > 3 {
            action.NewValue = commands[3]
        }

        handleAction(action)
    }
}