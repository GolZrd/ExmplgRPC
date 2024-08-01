package main

import (
	my_chat "ExmplgRPC/generated"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := my_chat.NewChatServiceClient(conn)

	response, err := c.Hi(context.Background(), &my_chat.ChatMessage{Text: "Привет тебе от клиента, чувак!"})
	if err != nil {
		panic(err)
	}
	log.Printf("Response from server: %s", response.Text)

}
