package servers

import (
	"context"
	pb "github.com/mpataki/playground-proto/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UserServer implements the UserService gRPC server
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// NewUserServer creates a new UserServer instance
func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
