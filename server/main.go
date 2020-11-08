package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "../logstream"

	"google.golang.org/grpc"
)

const (
	// defaultServerPort is the default port that used by the server.
	defaultServerPort = ":5050"

	// defaultServerEndpoint
	defaultServerEndpoint = "localhost"
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
	serverEndpoint := os.Getenv("ALPHA_SERVER_ENDPOINT")
	serverPort := os.Getenv("ALPHA_SERVER_PORT")

	if serverEndpoint == "" {
		serverEndpoint = defaultServerEndpoint
	}

	if serverPort == "" {
		serverPort = defaultServerPort
	}

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Alpha server started at %s%s\n", serverEndpoint, serverPort)
	}

	s := grpc.NewServer()
	pb.RegisterLogStreamerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
