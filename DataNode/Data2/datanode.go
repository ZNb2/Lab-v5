package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"strconv"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DataNode_name = "DataNode2"
//var Servidor_OMS = "localhost:50052"
var Servidor_OMS ="dist106.inf.santiago.usm.cl:50054"

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server)SayHello(ctx context.Context, in *pb.Message)(*pb.Message, error){

	log.Printf("Receive message body from client: %s", in.Body)
	directorio, _ := "DataNode/Data1/DATA.txt"
	
	if strings.Contains(in.Body, "::") {
		Escribir(in.Body, directorio)
		
	}else{

		content, err := os.ReadFile(directorio)
		if err != nil {
			log.Fatal(err)
		}

		lineas := strings.Split(string(content), "\n")

		for i := 0; i < len(lineas); i++ {
			
			split := strings.Split(lineas[i],"::")//id-nombre-apellido
			if split[0] == id {
				nombre_apellido := split[1]+"::"+split[2]
				ConexionGRPC(Servidor_OMS, nombre_apellido)
			}
		}
	}
	return &pb.Message{Body: "OK"}, nil
}

func ConexionGRPC(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println(Servidor, "not responding")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
}


func Escribir(mensaje string, nombreArchivo string) error {

	directorioActual, _ := os.Getwd()
	archivo, err := os.OpenFile(directorioActual+directorio, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer archivo.Close()

	// Escribimos el mensaje seguido de un salto de línea
	_, err = fmt.Fprintf(archivo, "%s\n", mensaje)
	if err != nil {
		return err
	}

	return nil
}


func main(){
	
	fmt.Println("Starting "+DataNode_name+" . . .")

	puerto := ":50054"
	lis, err := net.Listen("tcp", puerto)
	fmt.Printf("Escuchando %s\n", puerto)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	server := &Server{}
	pb.RegisterChatServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
