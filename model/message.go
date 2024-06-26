package model

import (
	"context"
	"encoding/json"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc/status"
)

type WsMessage struct {
	Event string   `json:"event"`
	Body  *Message `json:"body"`
	Query g.Map    `json:"query"`
	Tag   string   `json:"tag"`
	Uid   int64    `json:"_"`
}

func SyncWsMessage(ctx context.Context, event string, uid int64, data interface{}) {
	message := WsMessage{}
	message.Uid = uid
	message.Event = event
	message.Body = WrapMessage(data, nil)
	marshal, _ := json.Marshal(message)
	xetcd.Client.Put(ctx, "/sync/", string(marshal))
}

func WrapEventResponse(string string) string {
	return string + "_response"
}

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WrapMessage(data interface{}, err error) *Message {
	message := Message{}
	if err != nil {
		if s, ok := status.FromError(err); ok {
			message.Code = int(s.Code())
			message.Msg = s.Message()
		} else {
			message.Code = 1
			message.Msg = err.Error()
		}

	}
	message.Data = data
	return &message
}
