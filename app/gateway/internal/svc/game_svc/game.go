package game_svc

import (
	"context"
	"game/app/game/api/game/game"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	gameClient game.GameServiceClient
)

func GameClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("game_svc", grpcx.Balancer.WithRandom())
	gameClient = game.NewGameServiceClient(conn)
}
func ListBanner(ctx context.Context) (*game.BannerRes, error) {
	return gameClient.ListBanner(ctx, &game.BannerReq{})
}

func ListGameCategory(ctx context.Context) (*game.ListGameCategoryRes, error) {
	return gameClient.ListGameCategory(ctx, &game.ListGameCategoryReq{})
}
func ListGame(ctx context.Context, req *game.ListGameReq) (*game.ListGameRes, error) {
	return gameClient.ListGame(ctx, req)
}
