package main

import (
	"context"
	"log"
	"net"

	pb "go-grpc-rest-graphql/proto" // adjust path if needed

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Name: "Alice"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
