PROTO_DIR = api/protobuf
API_DIR = gen/proto
COLLECTOR_DIR = collector/internal
GATEWAY_DIR = gateway/internal

.PHONY: proto
proto:
	protoc --go_out=$(COLLECTOR_DIR)/$(API_DIR) --go_opt=paths=source_relative \
           --go-grpc_out=$(COLLECTOR_DIR)/$(API_DIR) --go-grpc_opt=paths=source_relative \
           -I=$(PROTO_DIR) \
           collector.proto; \
    protoc --go_out=$(GATEWAY_DIR)/$(API_DIR) --go_opt=paths=source_relative \
           --go-grpc_out=$(GATEWAY_DIR)/$(API_DIR) --go-grpc_opt=paths=source_relative \
           -I=$(PROTO_DIR) \
           collector.proto
