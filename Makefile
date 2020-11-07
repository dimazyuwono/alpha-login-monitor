build:
	go build -o client/main client/main.go
	go build -o server/main server/main.go


protobuff:
	bash logstream/generate.sh

run:
	go run server/main.go

