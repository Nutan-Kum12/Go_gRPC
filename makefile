run:
	go run server/main.go server/unary.go

client:
	go run client/main.go client/unary.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

run-all:
	start "" cmd /c "go run server/main.go server/unary.go"
	start "" cmd /c "go run client/main.go client/unary.go"

.PHONY: run client proto run-all