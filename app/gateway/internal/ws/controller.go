package ws

import (
	"context"
	"game/model"
)

type Controller interface {
	Controller(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error)
}
