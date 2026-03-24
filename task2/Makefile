PROTO_DIR = api/protobuf
API_DIR = internal/gen/proto

.PHONY: proto
proto:
	protoc --go_out=$(API_DIR) --go_opt=paths=source_relative \
           --go-grpc_out=$(API_DIR) --go-grpc_opt=paths=source_relative \
           -I=$(PROTO_DIR) \
           collector.proto

.PHONY: run
run:
	go run cmd/app/main.go
