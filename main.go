package main

import (
	my_chat "ExmplgRPC/generated"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, World!")

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	uss := &my_chat.MyBeautifulServer{}
	my_chat.RegisterChatServiceServer(grpcServer, uss)

	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
