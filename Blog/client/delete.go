package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"
)

func deleteblog(c pb.BlogServiceClient, ids string) {
	log.Println("delection have been invoked in client side")

	id := &pb.BlogId{Id: ids}

	c.DeleteBlog(context.Background(), id)
}
