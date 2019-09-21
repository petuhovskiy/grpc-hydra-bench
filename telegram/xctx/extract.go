package xctx

import (
	"context"

	"github.com/petuhovskiy/grpc-hydra-bench/telegram/desc"
)

func GetDesc(ctx context.Context) *desc.Main {
	return ctx.Value(Desc).(*desc.Main)
}
