package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/devishot/grpc-go-time_tracking/interface/api"
	"github.com/devishot/grpc-go-time_tracking/interface/handler"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler := &handler.Server{}
	grpcServer := grpc.NewServer()

	api.RegisterTimeTrackingServer(grpcServer, handler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
