package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrchestratorServiceServer
}

func (s *server) GetUserByName(ctx context.Context, in *pb.UserName) (*pb.User, error) {
	return nil, errors.New("not implemented yet. Souvik will implement me")
}

func main() {
	listner, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Can not bind server to port 3000: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrchestratorServiceServer(grpcServer, &server{})

	log.Printf("Server started on port 3000")
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalf("Error occured in serving grpc server %v", err)
	}
}
