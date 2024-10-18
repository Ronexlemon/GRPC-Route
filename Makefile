gen:
	@protoc \
	 --proto_path=protobuf "protobuf/route.proto" \
	 --go_out=services/genproto/route  --go_opt=paths=source_relative \
	 --go-grpc_out=services/genproto/route  --go-grpc_opt=paths=source_relative