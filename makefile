run:
	go run server/main.go server/unary.go server/server_stream.go server/client_stream.go server/bi_stream.go

client:
	go run client/main.go client/unary.go client/server_stream.go client/client_stream.go client/bi_stream.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

.PHONY: run client proto run-all