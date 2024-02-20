generate-go-grpc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/invoice.proto
run-server:
	go run ./cmd/server
