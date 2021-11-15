package main

import (
	"context"
	"fmt"
	"github.com/Souvikns/orchestrator-service/constants"
	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost" + constants.ORC_PORT1

func main() {
	fmt.Print("Enter username \n > ")
	var username string
	fmt.Scan(&username)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

	defer conn.Close()

	client := pb.NewOrchestratorServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	response, err := client.GetUserByName(ctx, &pb.UserName{Name: username})
	if err != nil {
		log.Fatalf("failed to get response %v", err)
	} else {
		fmt.Println("Name: " + response.Name)
		fmt.Println("Class: " + response.Class)
		fmt.Println("Roll: ", response.Roll)
	}
}
