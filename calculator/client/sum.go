package main

import (
	pb "MyCal/calculator/proto"
	"context"
	"log"
)

func Sum(c pb.CalculatorServiceClient) {
	res,err:=c.Sum(context.Background(),&pb.SumRequest{
		Number1: 2,
		Number2: 3,
	})
	if err != nil {
	log.Fatal(err)
	}
	log.Println("addtion")
	log.Println(res.Result)
}
