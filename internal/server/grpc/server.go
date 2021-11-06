package main

import (
	"context"
	"fmt"
	"log"
	"net"

	googleGRPC "google.golang.org/grpc"

	userGRPC "github.com/rodkevich/gmp-tickets/internal/grpcz/user"
)

// GRPCServer base type
type GRPCServer struct {
	userGRPC.UnimplementedUserServiceServer
}

func (s *GRPCServer) CreateUser(ctx context.Context, request *userGRPC.CreateUserRequest) (*userGRPC.CreateUserResponse, error) {

	return &userGRPC.CreateUserResponse{Token: "may be sessions ... ?"}, nil
}

func (s *GRPCServer) ReadUser(ctx context.Context, request *userGRPC.ReadUserRequest) (*userGRPC.ReadUserResponse, error) {

	return &userGRPC.ReadUserResponse{}, nil
}

func (s *GRPCServer) UpdateUser(ctx context.Context, request *userGRPC.UpdateUserRequest) (*userGRPC.ReadUserResponse, error) {
	return &userGRPC.ReadUserResponse{User: nil}, nil
}

func (s *GRPCServer) DeleteUser(ctx context.Context, request *userGRPC.DeleteUserRequest) (*userGRPC.DeleteUserResponse, error) {

	return &userGRPC.DeleteUserResponse{}, nil
}

func (s *GRPCServer) Login(ctx context.Context, request *userGRPC.LoginRequest) (*userGRPC.LoginResponse, error) {
	return &userGRPC.LoginResponse{Token: "may be sessions ... ?"}, nil
}

func (s *GRPCServer) Logout(ctx context.Context, request *userGRPC.LogoutRequest) (*userGRPC.LogoutResponse, error) {

	return &userGRPC.LogoutResponse{}, nil // may be outdated token ?
}

func (s *GRPCServer) CreateProfile(ctx context.Context, request *userGRPC.CreateProfileRequest) (*userGRPC.CreateProfileResponse, error) {

	return &userGRPC.CreateProfileResponse{}, nil
}

func (s *GRPCServer) ReadProfile(ctx context.Context, request *userGRPC.ReadProfileRequest) (*userGRPC.ReadProfileResponse, error) {

	return &userGRPC.ReadProfileResponse{}, nil
}

func (s *GRPCServer) UpdateProfile(ctx context.Context, request *userGRPC.UpdateProfileRequest) (*userGRPC.ReadProfileResponse, error) {

	return &userGRPC.ReadProfileResponse{}, nil
}

func (s *GRPCServer) DeleteProfile(ctx context.Context, request *userGRPC.DeleteProfileRequest) (*userGRPC.DeleteProfileResponse, error) {

	return &userGRPC.DeleteProfileResponse{}, nil
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
	userGRPC.RegisterUserServiceServer(s, grpcUsersServer)

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
