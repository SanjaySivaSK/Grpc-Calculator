package main

import (
	pb "MyCal/calculator/proto"
	"context"
	"log"
)

func Div(c pb.CalculatorServiceClient) {
	res,err:=c.Div(context.Background(),&pb.DivRequest{
		Number1: 10,
		Number2: 2,
	})
	if err != nil {
	log.Fatal(err)
	}
	log.Println("division")
	log.Println(res.Quotient)
	log.Println(res.Reminder)
}
