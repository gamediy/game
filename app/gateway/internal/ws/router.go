package ws

import (
	"context"
	"encoding/json"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	Error     = "error"
	Login     = "login"
	Join      = "join"
	Quit      = "quit"
	Heartbeat = "heartbeat"
	Enter     = "enter"
)

var (
	controller map[string]Controller
)

func init() {
	controller = make(map[string]Controller)
	controller["login"] = LoginController{}
	controller["heartbeat"] = HeartbeatController{}
}
func Router(ctx context.Context, client *Client, msg []byte) {
	message := model.Message{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	c, ok := controller[message.Event]
	if !ok {
		return
	}
	m, e := c.Controller(client, &message)
	if e != nil {
		g.Log().Error(ctx, e)
		return
	}
	if m != nil {
		client.WriteChn <- m
	}

}
