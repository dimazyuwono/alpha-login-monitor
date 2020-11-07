# Alpha Login Monitor

A simple client server based login monitoring for linux based operating system.

# Requirements

Please ensure thease application are installed in the machine:
- go
- make

# Guide

```bash
# First We'll need to build the application 
make buld

# Then to run the server
make run

```

Some parameters can be modified by importing the values into system environment variables.

| Variable Name  	        | Descriptin                        |
|---	                    |---	                            |
| ALPHA_AUTH_LOG_FILE_PATH  | path of the auth log that         |
| ALPHA_SERVER_ENDPOINT  	| Target endpoint to stream the log |
| ALPHA_SERVER_PORT  	    | Target port to stream the log     |
