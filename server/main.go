package main

import (
	"context"
	"log"
	"net"

	pb "../logstream"

	"google.golang.org/grpc"
)

const (
	// defaultServerPort is the default port that used by the server.
	defaultServerPort = ":5050"
)

// server is used to implement logstream.UnimplementedLogStreamerServer.
type server struct {
	pb.UnimplementedLogStreamerServer
}

// StreamLog implements logstream.LogStreamerServer
func (s *server) StreamLog(ctx context.Context, in *pb.LogStreamRequest) (*pb.LogStreamResponse, error) {
	log.Printf("%s had %d login attempt", in.GetHostname(), in.GetAttemp())
	return &pb.LogStreamResponse{Message: "Received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLogStreamerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
