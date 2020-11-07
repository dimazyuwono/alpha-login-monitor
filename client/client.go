package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

const (
	// authLogFilePath is the value of default OS auth log path.
	defaultAuthLogFilePath = "/var/log/auth.log"

	// defaultServerEndpoint is the default endpoint where the server hosted.
	defaultServerEndpoint = "127.0.0.1"

	// defaultServerPort is the default port that used to connect to server.
	defaultServerPort = "5050"
)

func main() {
	authLogFilePath := os.Getenv("ALPHA_AUTH_LOG_FILE_PATH")
	serverEndpoint := os.Getenv("ALPHA_SERVER_ENDPOINT")
	serverPort := os.Getenv("ALPHA_SERVER_PORT")

	if authLogFilePath == "" {
		authLogFilePath = defaultAuthLogFilePath
	}

	authLogStream, err := tail.TailFile(authLogFilePath, tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}

	for line := range authLogStream.Lines {
		fmt.Println(line.Text)
	}
}
