# Alpha Login Monitor

A simple client server based ssh login monitoring for linux based operating system.

# Requirements

Please ensure thease application are installed in the machine:
- go
- make

# Basic operation
## Build and Run
This repository contain simple Makefile that can help to simplify the build and run of the application.

```bash
# First We'll need to build the application 
make buld

# Then to run the server
make run
```
## Configuration

Some parameters can be modified by importing the values into system environment variables.

| Variable Name  	        | Descriptin                        |
|---	                    |---	                            |
| ALPHA_AUTH_LOG_FILE_PATH  | path of the auth log that         |
| ALPHA_SERVER_ENDPOINT  	| Target endpoint to stream the log |
| ALPHA_SERVER_PORT  	    | Target port to stream the log     |

# Todo

This application is not complete yet, there are a few tweaks that can be made:

- Implement limitter in sending the stream to the server.
- Instead creating 2 binaries for server and clien, We can create the application as one binary to simplify the build. 