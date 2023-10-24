package logic

import (
	"context"
	"game/app/user/api/user/user"
	"game/consts/user_event"
	"game/model"
)

func Test123(ctx context.Context) {
	reply := user.WalletReply{}
	reply.Balance = 3

	model.SyncWsMessage(ctx, model.WrapEventResponse(user_event.Wallet), 0, reply)

}
