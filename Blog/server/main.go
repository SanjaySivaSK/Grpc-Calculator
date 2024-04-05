package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50053"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("Blogdb").Collection("Blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening On", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
