package main

import (
	"fmt"
	"log"

	pb "MyCal/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50052"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hiiiiiiiiiiiiiiiih")
	
	c := pb.NewCalculatorServiceClient(conn)

	// Sum(c)
	// di(c)
	Primes(c)
	// Doaverage(c)
	// Domax(c)
	



	defer conn.Close()
}


