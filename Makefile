gen:
	protoc -I=. --go_out=. --go-grpc_out=. internal/proto/user/v1/user.proto
