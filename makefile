run:
	go run server/main.go

client:
	go run client/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto
	
run-all:
	start "" cmd /c "go run server/main.go"
	start "" cmd /c "go run client/main.go"

.PHONY: run client proto run-all