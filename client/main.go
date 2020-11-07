package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "../logstream"

	"github.com/hpcloud/tail"
	"google.golang.org/grpc"
)

const (
	// authLogFilePath is the value of default OS auth log path.
	defaultAuthLogFilePath = "/var/log/auth.log"

	// defaultServerEndpoint is the default endpoint where the server hosted.
	defaultServerEndpoint = "127.0.0.1"

	// defaultServerPort is the default port that used to connect to server.
	defaultServerPort = ":5050"

	// defaultHostname is the default hostname value that will be use to identify the client.
	defaultHostname = "localhost"

	// defaultLoginAttemp is the default login attempt value that will be sent to the server
	defaultLoginAttempt int32 = 0
)

func main() {
	// Read a environment variable to inject configuration

	authLogFilePath := os.Getenv("ALPHA_AUTH_LOG_FILE_PATH")
	serverEndpoint := os.Getenv("ALPHA_SERVER_ENDPOINT")
	serverPort := os.Getenv("ALPHA_SERVER_PORT")

	if authLogFilePath == "" {
		authLogFilePath = defaultAuthLogFilePath
	}

	if serverEndpoint == "" {
		serverEndpoint = defaultServerEndpoint
	}

	if serverPort == "" {
		serverPort = defaultServerPort
	}

	authLogStream, err := tail.TailFile(authLogFilePath, tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}

	for line := range authLogStream.Lines {
		fmt.Println(line.Text)

		// Set up a connection to the server.
		conn, err := grpc.Dial(serverEndpoint+serverPort, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewLogStreamerClient(conn)

		// Contact the server and print out its response.
		thisHostname := defaultHostname
		thisLoginAttemp := defaultLoginAttempt

		if len(os.Args) > 1 {
			thisHostname = os.Args[1]
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.StreamLog(ctx, &pb.LogStreamRequest{Hostname: thisHostname, Attemp: thisLoginAttemp})

		if err != nil {
			log.Fatalf("could not stream: %v", err)
		}
		log.Printf("Response: %s", r.GetMessage())
	}
}
