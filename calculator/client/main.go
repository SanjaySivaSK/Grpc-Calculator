package main

import (
	"fmt"
	"log"

	pb "MyCal/calculator/proto"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50053"

func main() {

	// opts := []grpc.DialOption{}

	// tls := true
	// certfile := "ssl/ca.crt"

	// if tls {
	// 	creds, err := credentials.NewClientTLSFromFile(certfile, "")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// }

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hiiiiiiiiiiiiiiiih")

	c := pb.NewCalculatorServiceClient(conn)

	Sum(c)

	// defer conn.Close()
} 