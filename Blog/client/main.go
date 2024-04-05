package main

import (
	"log"

	pb "MyCal/Blog/proto"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50053"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	ReadBlog(c, id)
	UpdateBlog(c, id)
	listofblog(c)
    deleteblog(c,"660d02994afa4a0989e24ffa")
	defer conn.Close()
}
