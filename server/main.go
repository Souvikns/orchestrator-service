package main

import (
	"context"
	"github.com/Souvikns/orchestrator-service/constants"
	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func callOrc2(username string) (*pb.User, error) {
	var address = "localhost" + constants.ORC_PORT2
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

	defer conn.Close()
	client := pb.NewOrchestrator2ServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	respose, err := client.GetUser(ctx, &pb.UserName{Name: username})
	if err != nil {
		return nil, err
	}

	return respose, err
}

type server struct {
	pb.UnimplementedOrchestratorServiceServer
}

func (s *server) GetUserByName(ctx context.Context, request *pb.UserName) (*pb.User, error) {
	user, err := callOrc2(request.Name)
	if err != nil {
		return nil, err
	}
	return user, err
}

func main() {
	listner, err := net.Listen("tcp", constants.ORC_PORT1)
	if err != nil {
		log.Fatalf("Can not bind server to port 3000: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrchestratorServiceServer(grpcServer, &server{})

	log.Printf("Server started on port " + constants.ORC_PORT1)
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalf("Error occured in serving grpc server %v", err)
	}
}
