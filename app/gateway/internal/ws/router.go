package ws

import (
	"context"
	"encoding/json"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

// event path

func MakeController() {

	ControllerC = make(map[string]Controller, 100)
	UserControllerInit()
}
func Router(ctx context.Context, client *Client, msg []byte) {

	message := model.WsMessage{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	c, ok := ControllerC[message.Event]
	if !ok {
		return
	}
	m, e := c.Controller(ctx, client, &message)
	if e != nil {
		g.Log().Error(ctx, e)
		return
	}
	if m != nil {
		client.WriteChn <- m
	}

}
