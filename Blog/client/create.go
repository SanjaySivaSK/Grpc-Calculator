package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {

	log.Println("client side create blog")

	blog := &pb.Blog{
		AuthorId: "SANJAY",
		Title:    "KUMAR",
		Content:  "SK",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Id)
	return res.Id
}
