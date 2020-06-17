package main

import (
	"context"
	"log"
	"net"

	"github.com/Dariellambert/mygolang-webapp/proto"
	"google.golang.org/grpc"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	return &proto.AuthResponse{}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
	listener, err := net.Listen("tcp", ":50000")
	if err != nill {
		log.Fatal("Error creating listener: ", err.Error())
	}
	server.Serve(listener)
}
