package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-rest-graphql/proto" // adjust path if needed

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from gRPC server: %s\n", resp.Name)
}
