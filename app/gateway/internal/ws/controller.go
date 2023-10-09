package ws

import (
	"context"
	"game/model"
)

var (
	ControllerC map[string]Controller
)

type Controller interface {
	Controller(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error)
}
