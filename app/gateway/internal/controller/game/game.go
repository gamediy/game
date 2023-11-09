package game

import (
	"context"
	"game/app/game/api/game/game"
	"game/app/gateway/internal/controller"
	"game/app/gateway/internal/svc/game_svc"
	"game/app/gateway/internal/ws"
	"game/consts/event/game_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func ControllerInit() {
	controller.Ctrl[game_event.ListBanner] = listBanner
	controller.Ctrl[game_event.ListGameCategory] = listGameCategory
	controller.Ctrl[game_event.ListGame] = listGame
}

func listGame(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	req := &game.ListGameReq{}
	req.Code = gconv.String(query["code"])
	res, err := game_svc.ListGame(ctx, req)
	return &model.WsMessage{
		Event: model.WrapEventResponse(game_event.ListGame),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listGameCategory(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := game_svc.ListGameCategory(ctx)
	return &model.WsMessage{
		Event: model.WrapEventResponse(game_event.ListGameCategory),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listBanner(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := game_svc.ListBanner(ctx)
	return &model.WsMessage{
		Event: model.WrapEventResponse(game_event.ListBanner),
		Body:  model.WrapMessage(res, err),
	}, nil
}
