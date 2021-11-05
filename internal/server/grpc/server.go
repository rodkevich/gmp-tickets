package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	userGRPC "github.com/rodkevich/gmp-tickets/internal/user/grpc"
)

// GRPCServer base type
type GRPCServer struct {
	userGRPC.UnimplementedUserServiceServer

	// db
}

// Run start a server
func (s *GRPCServer) Run(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}
func main() {
	startApp("127.0.0.1:9090")
}

func startApp(address string) {
	s := grpc.NewServer()

	grpcUsersServer := &GRPCServer{}
	// init fake database
	// if err := grpcUsersServer.InitSomeDb(); err != nil {
	// 	log.Fatal(err)
	// }
	// register services
	userGRPC.RegisterUserServiceServer(s, grpcUsersServer)

	usersInstance, err := grpcUsersServer.Run(address)
	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	// serve grpc
	if err := s.Serve(usersInstance); err != nil {
		log.Fatal(err)
	}
}
