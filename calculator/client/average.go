package main

import (
	pb "MyCal/calculator/proto"
	"context"
	"log"
)

func Doaverage(c pb.CalculatorServiceClient) {

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	// reqs := []*pb.AverageRequest{
	// 	{Number: 1},
	// 	{Number: 2},
	// 	{Number: 2},
	// 	{Number: 3},
	// 	{Number: 1},
	// 	{Number: 6},
	// 
	reqs:=[]int64{1,3,3,4,4,4,4,4}

	for _, req := range reqs {
		stream.Send(&pb.AverageRequest{
			Number: req,
		})

	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)

}
