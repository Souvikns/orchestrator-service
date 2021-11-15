package main

import (
	"context"
	"github.com/Souvikns/orchestrator-service/constants"
	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func loadMockData(username string) (*pb.User, error) {
	var address = "localhost" + constants.DUMMY_DATA_PORT
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}
	defer conn.Close()

	client := pb.NewMockDataServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	response, err := client.GetMockUserData(ctx, &pb.UserName{Name: username})
	if err != nil {
		return nil, err
	}

	return response, nil
}

type server struct {
	pb.UnimplementedOrchestrator2ServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserName) (*pb.User, error) {
	user, err := loadMockData(req.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	app, err := net.Listen("tcp", constants.ORC_PORT2)
	if err != nil {
		log.Fatalf("Can not bind to server to port 9000 %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrchestrator2ServiceServer(grpcServer, &server{})

	log.Printf("Server started on port " + constants.ORC_PORT2)
	err = grpcServer.Serve(app)
	if err != nil {
		log.Fatalf("Error occured in serving gcp server %v", err)
	}
}
