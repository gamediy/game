package logic

import (
	"context"
	"game/app/user/api/user/user"
	"game/consts/event/user_event/wallet_event"
	"game/model"
)

func Test123(ctx context.Context) {
	reply := user.WalletReply{}
	reply.Balance = 3

	model.SyncWsMessage(ctx, model.WrapEventResponse(wallet_event.Wallet), 0, reply)

}
