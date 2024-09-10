run-client:
	@go run cmd/client/*.go

run-restaurant:
	@go run cmd/restaurant/*.go

run-server:
	@go run cmd/server/*.go

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=internal/gen --go_opt=paths=source_relative \
  	--go-grpc_out=internal/gen --go-grpc_opt=paths=source_relative		