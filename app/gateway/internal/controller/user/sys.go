package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/sys"
	"game/consts/event/user_event/sys_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func getAnnouncement(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.GetAnnouncement(ctx, &sys.AnnouncementReq{})
	return &model.WsMessage{
		Event: model.WrapEventResponse(sys_event.GetAnnouncement),
		Body:  model.WrapMessage(res, err),
	}, nil
}
