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
	Inicial = map[string]string{
		  "A": "DataNode 1" ,"B": "DataNode 1" ,"C": "DataNode 1" ,"D": "DataNode 1" ,
		  "E": "DataNode 1" ,"F": "DataNode 1" ,"G": "DataNode 1" ,"H": "DataNode 1" ,
		  "I": "DataNode 1" ,"J": "DataNode 1" ,"K": "DataNode 1" ,"L": "DataNode 1" ,
		  "M": "DataNode 1" ,"N": "DataNode 2" ,"O": "DataNode 2" ,"P": "DataNode 2" ,
		  "Q": "DataNode 2" ,"R": "DataNode 2" ,"S": "DataNode 2" ,"T": "DataNode 2" ,
		  "U": "DataNode 2" ,"V": "DataNode 2" ,"W": "DataNode 2" ,"X": "DataNode 2" ,
		  "Y": "DataNode 2" ,"Z": "DataNode 2" ,
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
		p1 = Servidores[string(pi[len(p1)-1]) + ":50052"]
		mensaje := strings.Split(in.Body, "--")
		//Datanode := Inicial[string(mensaje[1][0])]
		//log.Printf("%s, %s", Datanode, Servidores[p1])
		
		id++
		//Añadir al txt: id, datanode, estado
		msj_datanode := strconv.Itoa(id) +"-"+ mensaje[0] +"-"+ mensaje[1]
		log.Printf("Solicitud de %s recibida, mensaje enviado: %s", p1, msj_datanode)
		//ConexionGRPC(Datanode, msj_datanode)
		
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



func main() {
	
	server_name := "OMS"
	fmt.Println("Starting "+server_name+" . . .")
	
	go Escuchar(":50052")
	Escuchar(":50053")
	
}



