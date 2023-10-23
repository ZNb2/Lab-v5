package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var server_name = "ONU"
//var Servidor_OMS = "localhost:50052"
var Servidor_OMS ="dist106.inf.santiago.usm.cl:50053"

func ConexionGRPC(mensaje string ){
	
	conn, err := grpc.Dial(Servidor_OMS, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println("Server OMS not responding ")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	log.Println("Estado enviado:", strings.Replace(mensaje, "--", " ", -1))
}

var (
	Estados = map[string]string{
		"I": "Infectado",
		"M": "Muerto",
	}
)

func main() {

	fmt.Println("Starting " + server_name + " . . .\n")

	for {
	fmt.Print("Seleccione Tipo (I/M): ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\r\n")
	input = strings.ToUpper(input)
	if input == "I" || input =="M" {
		//fmt.Println(input)
		ConexionGRPC(Estados[input])
	}else{
		fmt.Println("\nComando no reconocido!\n")
	}
	//fmt.Println(input)

	}

}