# Example

To run this demo / example you will need to ensure that your local machine have `go`and `make` installed in the system.

```bash
# First of all you will need to clone this repository
$ git clone https://github.com/syhrz/alpha-login-monitor
$ cd alpha-login-monitor

# build the binaries
$ make build

# It will generate 2 files inside
# client/main and server/main

# run the server 
$ make run

# The server will run at the localhost with the default 5050 port 
# For this demo We will run both server and client in one machine

# Please open a new terminal session

# Before starting the client, We will need to setup a few parameters
# copy the example_auth.log to /tmp
$ cp example/example_auth.log /tmp

# Let's use the example auth.log
$ export ALPHA_AUTH_LOG_FILE_PATH="/tmp/example_auth.log"

# Then start the client by running
$ ./client/main

# The client will start reading the defined auth log path and stream it to the server.
# To stop the application go to both session and press ctrl+c to terminate the process
```
