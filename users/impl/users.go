package impl

import (
	"context"

	"github.com/petuhovskiy/grpc-hydra-bench/users/pb"
)

type UsersServer struct{}

func (s *UsersServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	panic("implement me")
}
