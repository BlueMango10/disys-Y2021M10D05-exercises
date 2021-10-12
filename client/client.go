package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"ithub.com/BlueMango10/disys-Y2021M10D05-exercises/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedGetXXXServer
}

func (s *Server) GetTime(ctx context.Context, in *time.GetXXXRequest) (*time.GetXXXReply, error) {
	fmt.Printf("Received XXX request")
	return &time.GetXXXReply{Reply: "Your reply here"}, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterXXXServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}