package main

import (
	"fmt"
	"log"
	"net"

	pb "MyCal/calculator/proto"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50053"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening On", addr)

	// tsl := true

	// opts := []grpc.ServerOption{}
	// if tsl {
	// 	certFile := "ssl/server.crt"
	// 	keyfile := "ssl/server.pem"

	// 	creds, err := credentials.NewServerTLSFromFile(certFile, keyfile)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	opts = append(opts, grpc.Creds(creds))

	// }

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)
	

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
