package game

import (
	"context"
	"game/app/gateway/internal/controller"
	"game/app/gateway/internal/svc/game_svc"
	"game/app/gateway/internal/ws"
	"game/consts/event/game_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func ControllerInit() {
	controller.Ctrl[game_event.ListBanner] = listBanner
}

func listBanner(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := game_svc.Service.ListBanner(ctx)
	return &model.WsMessage{
		Event: model.WrapEventResponse(game_event.ListBanner),
		Body:  model.WrapMessage(res, err),
	}, nil
}
