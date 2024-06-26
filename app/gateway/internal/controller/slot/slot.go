package slot

import (
	"context"
	"game/app/gateway/internal/controller"
	"game/app/gateway/internal/controller/user"
	"game/app/gateway/internal/svc/slot_svc"
	"game/app/gateway/internal/ws"
	"game/app/slot/api/slot/slot"
	"game/consts/event/slot_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func SlotControllerInit() {
	controller.Ctrl[slot_event.SlotSpin] = slotSpin
	controller.Ctrl[slot_event.SlotCheckWon] = slotCheckWon

}

func slotSpin(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {

	f := query["amount"].(float64)
	res, err := slot_svc.Service.Spin(ctx, &slot.SpinReq{
		Uid:      wsclient.UserInfo.Uid,
		Amount:   f,
		GameCode: gconv.Int32(query["gameCode"]),
	})
	message := model.WsMessage{
		Event: model.WrapEventResponse(slot_event.SlotSpin),
		Body:  model.WrapMessage(res, err),
	}
	if err == nil {
		user.SendWallet(ctx, wsclient)
	}
	return &message, err

}

func slotCheckWon(ctx context.Context, client *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := slot_svc.Service.CheckWon(ctx, &slot.CheckWonReq{
		OrderNo: gconv.Int64(query["orderNo"]),
		Uid:     client.UserInfo.Uid,
	})
	if res.Status == 2 {
		user.SendWallet(ctx, client)
	}
	message := model.WsMessage{
		Event: model.WrapEventResponse(slot_event.SlotCheckWon),
		Body:  model.WrapMessage(res, err),
	}
	return &message, err
}
