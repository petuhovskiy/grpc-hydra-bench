package impl

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/pb"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/users"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/xctx"
)

type AuthServer struct {
	users *users.Repo
}

func NewAuthServer(repo *users.Repo) *AuthServer {
	return &AuthServer{
		users: repo,
	}
}

func (s *AuthServer) GetUserInfo(ctx context.Context, _ *empty.Empty) (*pb.UserInfo, error) {
	token, err := xctx.GetInrospectedToken(ctx)
	if err != nil {
		return nil, err
	}

	spew.Dump(token)
	panic("implement me")
}
