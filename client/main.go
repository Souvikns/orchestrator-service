package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Souvikns/orchestrator-service/user"
	"google.golang.org/grpc"
)

const address = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

	defer conn.Close()

	client := pb.NewOrchestratorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetUserByName(ctx, &pb.UserName{Name: "Souvik"})
	if err != nil {
		log.Fatalf("failed to get response %v", err)
	} else {
		fmt.Println(response);
	}
}
