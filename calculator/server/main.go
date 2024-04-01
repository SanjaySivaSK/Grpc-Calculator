package main

import (
	"fmt"
	"log"
	"net"

	pb "MyCal/calculator/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50052"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening On" ,addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
