package main

import (
	"context"
	"fmt"
	"log"
	"net"

	googleGRPC "google.golang.org/grpc"

	"github.com/rodkevich/gmp-tickets/internal/user/grpc"
)

// GRPCServer base type
type GRPCServer struct {
	grpc.UnimplementedUserServiceServer
}

func (s *GRPCServer) CreateUser(ctx context.Context, request *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {

	return &grpc.CreateUserResponse{Token: "may be sessions ... ?"}, nil
}

func (s *GRPCServer) ReadUser(ctx context.Context, request *grpc.ReadUserRequest) (*grpc.ReadUserResponse, error) {

	return &grpc.ReadUserResponse{}, nil
}

func (s *GRPCServer) UpdateUser(ctx context.Context, request *grpc.UpdateUserRequest) (*grpc.ReadUserResponse, error) {
	return &grpc.ReadUserResponse{User: nil}, nil
}

func (s *GRPCServer) DeleteUser(ctx context.Context, request *grpc.DeleteUserRequest) (*grpc.DeleteUserResponse, error) {

	return &grpc.DeleteUserResponse{}, nil
}

func (s *GRPCServer) Login(ctx context.Context, request *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	return &grpc.LoginResponse{Token: "may be sessions ... ?"}, nil
}

func (s *GRPCServer) Logout(ctx context.Context, request *grpc.LogoutRequest) (*grpc.LogoutResponse, error) {

	return &grpc.LogoutResponse{}, nil // may be outdated token ?
}

func (s *GRPCServer) CreateProfile(ctx context.Context, request *grpc.CreateProfileRequest) (*grpc.CreateProfileResponse, error) {

	return &grpc.CreateProfileResponse{}, nil
}

func (s *GRPCServer) ReadProfile(ctx context.Context, request *grpc.ReadProfileRequest) (*grpc.ReadProfileResponse, error) {

	return &grpc.ReadProfileResponse{}, nil
}

func (s *GRPCServer) UpdateProfile(ctx context.Context, request *grpc.UpdateProfileRequest) (*grpc.ReadProfileResponse, error) {

	return &grpc.ReadProfileResponse{}, nil
}

func (s *GRPCServer) DeleteProfile(ctx context.Context, request *grpc.DeleteProfileRequest) (*grpc.DeleteProfileResponse, error) {

	return &grpc.DeleteProfileResponse{}, nil
}

// Run start a server
func (s *GRPCServer) Run(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}
func main() {
	startApp("127.0.0.1:12312")
}

func startApp(address string) {
	s := googleGRPC.NewServer()

	grpcUsersServer := &GRPCServer{}
	// init fake database
	// if err := grpcUsersServer.InitSomeDb(); err != nil {
	// 	log.Fatal(err)
	// }
	// register services
	grpc.RegisterUserServiceServer(s, grpcUsersServer)

	usersInstance, err := grpcUsersServer.Run(address)
	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	// serve googleGRPC
	if err := s.Serve(usersInstance); err != nil {
		log.Fatal(err)
	}
}
