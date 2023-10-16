package ws

import (
	"context"
	"game/app/gateway/internal/svc/slot_svc"
	"game/app/slot/api/slot/slot"
	"game/consts/ws_consts"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func SlotControllerInit() {
	Ctrl[ws_consts.SlotSpin] = slotSpin
	Ctrl[ws_consts.SlotCheckWon] = slotCheckWon

}

func slotSpin(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error) {

	f := query["amount"].(float64)
	res, err := slot_svc.Service.Spin(ctx, &slot.SpinReq{
		Uid:      wsclient.UserInfo.Uid,
		Amount:   f,
		GameCode: gconv.Int32(query["gameCode"]),
	})
	message := model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.SlotSpin),
		Body:  model.WrapMessage(res, err),
	}
	if err == nil {
		SendWallet(ctx, wsclient)
	}
	return &message, err

}

func slotCheckWon(ctx context.Context, client *Client, query g.Map) (*model.WsMessage, error) {
	res, err := slot_svc.Service.CheckWon(ctx, &slot.CheckWonReq{
		OrderNo: gconv.Int64(query["orderNo"]),
		Uid:     client.UserInfo.Uid,
	})
	if res.Status == 2 {
		SendWallet(ctx, client)
	}
	message := model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.SlotCheckWon),
		Body:  model.WrapMessage(res, err),
	}
	return &message, err
}
