package game

import (
	"context"
	"game/app/game/api/game/game"
	"game/app/game/internal/service/gamesvc"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	game.UnimplementedGameServiceServer
}

func Register(s *grpcx.GrpcServer) {
	game.RegisterGameServiceServer(s.Server, &Controller{})
}

func (*Controller) ListBanner(ctx context.Context, req *game.BannerReq) (res *game.BannerRes, err error) {
	return gamesvc.ListBanner(ctx)
}

func (*Controller) ListGameCategory(ctx context.Context, req *game.ListGameCategoryReq) (res *game.ListGameCategoryRes, err error) {
	return gamesvc.ListGameCategory(ctx)
}

func (*Controller) ListGame(ctx context.Context, req *game.ListGameReq) (res *game.ListGameRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
