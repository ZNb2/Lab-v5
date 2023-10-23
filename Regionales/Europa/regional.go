package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var server_name = "Europa"
//var Servidor_OMS = "localhost:50052"
var Servidor_OMS ="dist105.inf.santiago.usm.cl:50052"

func ConexionGRPC(mensaje string ){
	
	conn, err := grpc.Dial(Servidor_OMS, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		log.Println("Estado enviado:", strings.Replace(mensaje, "--", " ", -1))
		_, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println("Server OMS not responding ")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
}

func ObtenerNombre() string{

	var rand_num int
	if len(Lista) == 0 {
		fmt.Println("\nNo hay mas nombres disponibles")
		os.Exit(0)
	}

	if len(Lista) == 1 {
		rand_num = 0
	}else{
		rand_num = rand.Intn(len(Lista)-1)
	}
	
	nombre := strings.Split(Lista[rand_num], " ")	
	Lista[rand_num] = Lista[len(Lista)-1]
	Lista = Lista[:len(Lista)-1]
	
	return nombre[0] + "--" + nombre[1]
}

func ObtenerStatus() string{

	if rand.Intn(100) > 55{
		return "muerto"
	}else {
		return "infectado"
	}
}

var Lista []string 

func main() {
		
	fmt.Println("Iniciando regional "+server_name+" . . .\n")
	rand.Seed(time.Now().UnixNano())
			
	directorioActual, _ := os.Getwd()
	content, err := os.ReadFile(directorioActual+"Regionales/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	Lista = strings.Split(string(content), "\n")
	Lista = Lista[:len(Lista)-1]

	//MANDAR 5 DATOS INICIALES
	for i := 0; i < 5; i++ {
		ConexionGRPC(ObtenerNombre() +"--"+ ObtenerStatus())
	}

	fmt.Println("\nSe mandaron 5 Nombres iniciales ...\nMandando datos cada 3 segundos ...\n")

	//MANDAR DATOS CADA 3 SEGUNDOS
	for{
		time.Sleep(3*time.Second)
		ConexionGRPC(ObtenerNombre() +"--"+ ObtenerStatus())
	}
}



