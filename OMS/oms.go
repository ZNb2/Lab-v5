package main

import (
	"context"
	"fmt"
	"log"
	"net"
	//"os"
	"strings"
	"strconv"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Servidores = map[string]string{
		"5:50052": "Australia",
		"6:50052": "Asia",
		"7:50052": "Europa",
		"8:50052": "Latinoamerica",
		
		"5:50053": "ONU",
		"7:50053": "DataNode 1",
		"8:50053": "DataNode 2",
	}
)

var (
	Node = map[string]string{
		"1": "dist107.inf.santiago.usm.cl:50053",
		"2": "dist108.inf.santiago.usm.cl:50053",
	}
)

var (
	Inicial = map[string]string{
		  "A": "1" ,"B": "1" ,"C": "1" ,"D": "1" ,
		  "E": "1" ,"F": "1" ,"G": "1" ,"H": "1" ,
		  "I": "1" ,"J": "1" ,"K": "1" ,"L": "1" ,
		  "M": "1" ,"N": "2" ,"O": "2" ,"P": "2" ,
		  "Q": "2" ,"R": "2" ,"S": "2" ,"T": "2" ,
		  "U": "2" ,"V": "2" ,"W": "2" ,"X": "2" ,
		  "Y": "2" ,"Z": "2" ,
	}
)

type Server struct{
	pb.UnimplementedChatServiceServer
}

var msj_data,id int
var Listado string

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	
	p, _ := peer.FromContext(ctx)
	p1 := strings.Split(p.Addr.String(), ":")[0]
		
	if strings.Contains(in.Body, "--"){
		//Mensaje de Contiente
		p1 = Servidores[string(p1[len(p1)-1]) + ":50052"]
		mensaje := strings.Split(in.Body, "--")
		datanode := Inicial[string(mensaje[1][0])]
		
		id++
		Escribir("DATA.txt", id+","+datanode+","+estado)
		msj_datanode := strconv.Itoa(id) +"-"+ mensaje[0] +"-"+ mensaje[1]
		log.Printf("Solicitud de %s recibida, mensaje enviado: %s", p1, msj_datanode)
		//ConexionGRPC(Node[datanode], msj_datanode)
		
	} else if strings.Contains(in.Body, "-"){
		//Mensaje de Datanode
		
		num := 10 //Obtener
		for i := 0; i < num; i++ {
			//ConexionGRPC(ONU, Listado)
		}
		
		for{
			if num == msj_data{
				msj_data = 0
				Listado = ""
		        }
		}
		
	} else {
		//Mensaje de ONU
		log.Printf("ONU")
		//ConexionGRPC(Datanode1, in.Body)
		//ConexionGRPC(Datanode2, in.Body)
		
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
		log.Println("Estado enviado:", strings.Replace(mensaje, "-", " ", -1))
		response, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if strings.Contains(response.Body, "-"){
			Listado += response.Body + "\n"
			msj_data++
		}
		if err != nil {
			log.Println(Servidor, "not responding")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}

}

func Escuchar(puerto string){
	
	lis_regional, err_regional:= net.Listen("tcp", puerto)
	fmt.Printf("Escuchando %s\n", puerto)
	if err_regional != nil {
		panic(err_regional)
	}

	grpcServer_regional := grpc.NewServer()
	server_regional := &Server{}
	pb.RegisterChatServiceServer(grpcServer_regional, server_regional)
	if err_regional := grpcServer_regional.Serve(lis_regional); err_regional != nil {
		panic(err_regional)
	}
}

func Escribir(mensaje string, nombreArchivo string) error {

	archivo, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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


func main() {
	
	server_name := "OMS"
	fmt.Println("Starting "+server_name+" . . .")
	
	go Escuchar(":50052")
	Escuchar(":50053")
	
}



