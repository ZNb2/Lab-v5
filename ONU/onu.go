package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo27-Laboratorio-2/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func ConexionGRPC(mensaje string ){
	
	//Uno de estos debe cambiar quizas por "regional:50052" ya que estara en la misma VM que el central
	//host :="localhost"
	var puerto, nombre, host string
	host="dist105.inf.santiago.usm.cl"
	puerto ="50053"
	nombre ="OMS"
	
	log.Println("Connecting to server "+nombre+": "+host+":"+puerto+". . .")
	conn, err := grpc.Dial(host+":"+puerto,grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	log.Printf("Esperando\n")
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	for {
		log.Println("Sending message to server "+nombre+": "+mensaje)
		response, err := c.OnuToOms(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println("Server "+nombre+" not responding: ")
			log.Println("Trying again in 10 seconds. . .")
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Response from server "+nombre+": "+"\n%s\n", response.Body)
		break
	}
}



var server_name string
func main() {
	server_name="ONU"
	log.Println("Starting " + server_name + " . . .\n")

	for {
	log.Print("Seleccione Tipo (I/M): ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.ToUpper(input)
	if input == "I" || input =="M" {
		log.Println(input)
		ConexionGRPC(input)
	}else{
		log.Println("\nComando no reconocido!\n")
	}
	//fmt.Println(input)

	}

}
