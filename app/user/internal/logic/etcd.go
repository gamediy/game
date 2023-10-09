package logic

import (
	"context"
	"game/app/user/api/user/user"
	"game/consts/ws_consts"
	"game/model"
)

func Test123(ctx context.Context) {
	reply := user.WalletReply{}
	reply.Balance = 3

	model.SyncWsMessage(ctx, model.WrapEventResponse(ws_consts.Wallet), 0, reply)

}
