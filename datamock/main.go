package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"github.com/Souvikns/orchestrator-service/constants"
	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMockDataServiceServer
}

func (server *server) GetMockUserData(ctx context.Context, request *pb.UserName) (*pb.User, error) {
	if len(request.Name) < 6 {
		return nil, errors.New("username must be more than 6 charachters")
	}

	return &pb.User{Name: request.Name, Class: fmt.Sprint(len(request.Name)), Roll: int64(len(request.Name)) * 10}, nil
}

func main() {
	listner, err := net.Listen("tcp", constants.DUMMY_DATA_PORT)
	if err != nil {
		log.Fatalf("Can not bind server to port " + constants.DUMMY_DATA_PORT)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMockDataServiceServer(grpcServer, &server{})
	log.Printf("Server started on port " + constants.DUMMY_DATA_PORT)
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalf("Error occured in serving grpc server %v", err)
	}
}
