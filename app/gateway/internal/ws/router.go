package ws

import (
	"context"
	"encoding/json"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// event path
var (
	Ctrl map[string]func(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error)
)

func MakeController() {

	Ctrl = make(map[string]func(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error), 100)
	UserControllerInit()
	SlotControllerInit()
}
func Router(ctx context.Context, client *Client, msg []byte) {

	message := model.WsMessage{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	c, ok := Ctrl[message.Event]
	if !ok {
		return
	}
	m2 := gconv.Map(message.Query)
	m, e := c(ctx, client, m2)
	if e != nil {
		g.Log().Error(ctx, e)
	}
	if m != nil {
		client.WriteChn <- m
	}

}
