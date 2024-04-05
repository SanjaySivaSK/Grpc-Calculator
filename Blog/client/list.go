package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listofblog(c pb.BlogServiceClient) {

	log.Println("client side call have been invoked (list of blog)")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err, " the error occurs in stream of list of blog")
	}

	for {
		datas, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err, " the error occure in the receive of list of blog")
			break
		}

		log.Println(datas)

	}

}
