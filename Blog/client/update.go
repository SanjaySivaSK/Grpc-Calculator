package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"
)

func UpdateBlog(c pb.BlogServiceClient, id string) {

	log.Println("Update blog called")

	data := &pb.Blog{
		Id:       id,
		AuthorId: "sssss",
		Title:    "ssss",
		Content:  "sss",
	}

	_, err := c.UpdateBlog(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("updated ")

}
