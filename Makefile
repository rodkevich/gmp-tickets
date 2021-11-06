gen:
	protoc -I=. --go_out=. --go-grpc_out=. internal/proto/user/v1/user.proto

ticket:
	 curl -d "@examples/createTicket.json" POST 0.0.0.0:12300/api/v1/tickets -v -H "Content-Type: application/json"
