package main

import (
	"context"
	"log"
	"os"
	"regexp"
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
	defaultClientHostname = "localhost"

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

	thisClientHostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		thisClientHostname = defaultClientHostname
	}

	log.Printf("Client started at %s reading from %s", thisClientHostname, authLogFilePath)

	numberOfLoginAttemp := make(chan int32)

	go initializeAuthLogStream(authLogFilePath, numberOfLoginAttemp)

	for thisNumberOfLoginAttemp := range numberOfLoginAttemp {
		// thisNumberOfLoginAttemp := <-numberOfLoginAttemp
		// log.Printf("%d", thisNumberOfLoginAttemp)
		sentMetricstoServer(serverEndpoint, serverPort, thisClientHostname, thisNumberOfLoginAttemp)
	}
}

func sentMetricstoServer(serverEndpoint string, serverPort string, clientHostname string, numberOfLoginAttemp int32) {
	// Set up a connection to the server.

	log.Printf("Trying to start connection to %s%s", serverEndpoint, serverPort)

	conn, err := grpc.Dial(serverEndpoint+serverPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Printf("Client connected to %s%s", serverEndpoint, serverPort)
	}
	defer conn.Close()
	c := pb.NewLogStreamerClient(conn)

	log.Printf("Sending the message to %s%s", serverEndpoint, serverPort)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StreamLog(ctx, &pb.LogStreamRequest{Hostname: clientHostname, Attemp: numberOfLoginAttemp})

	if err != nil {
		log.Fatalf("could not reach server: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}

func initializeAuthLogStream(authLogFilePath string, numberOfLoginAttemp chan int32) {
	authLogStream, err := tail.TailFile(authLogFilePath, tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}

	thisNumberOfLoginAttemp := defaultLoginAttempt

	for line := range authLogStream.Lines {
		getLoginAttemp, err := regexp.MatchString("Accepted*", line.Text)

		if err != nil {
			log.Fatal(err)
		} else if getLoginAttemp {
			log.Printf("User login detected")
			thisNumberOfLoginAttemp++
		}

		log.Printf("Total number %d", thisNumberOfLoginAttemp)
		numberOfLoginAttemp <- thisNumberOfLoginAttemp
	}
}
