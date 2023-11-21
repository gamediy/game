package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/mailbox"
	"game/consts/event/user_event/mailbox_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func countMailBoxTotal(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	read, err := user_svc.Service.CountMailBoxUnRead(ctx, wsclient.UserInfo.Uid)
	return &model.WsMessage{
		Event: model.WrapEventResponse(mailbox_event.MailBoxTotal),
		Body:  model.WrapMessage(read, err),
	}, nil
}

func listMailBox(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	req := mailbox.ListMailBoxReq{}
	if query == nil {
		query = g.Map{"size": 10, "page": 1}
	}
	req.Size = gconv.Int64(query["size"])
	req.Page = gconv.Int64(query["page"])
	req.Read = gconv.String(query["read"])
	req.Receiver = gconv.String(wsclient.UserInfo.Uid)
	req.Type = gconv.String(query["implement"])
	res, err := user_svc.Service.ListMailBox(ctx, &req)
	return &model.WsMessage{
		Event: model.WrapEventResponse(mailbox_event.ListMailBox),
		Body:  model.WrapMessage(res, err),
	}, nil
}
