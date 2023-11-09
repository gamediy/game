package game_svc

import (
	"context"
	"game/app/game/api/game/game"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	Service    = &gameSvc{}
	gameClient game.GameServiceClient
)

type gameSvc struct{}

func GameClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("game_svc", grpcx.Balancer.WithRandom())
	gameClient = game.NewGameServiceClient(conn)
}
func (gameSvc) ListBanner(ctx context.Context) (*game.BannerRes, error) {
	return gameClient.ListBanner(ctx, &game.BannerReq{})
}
