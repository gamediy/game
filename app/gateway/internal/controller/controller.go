package controller

import (
	"context"
	"game/app/gateway/internal/ws"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	Ctrl = make(map[string]func(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error), 100)
}

// event path
var (
	Ctrl map[string]func(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error)
)
