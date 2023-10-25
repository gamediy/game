package wallet_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/db/dao"
)

func Wallet(ctx context.Context, req *user.WalletRequest) (reply *user.WalletReply, err error) {
	reply = &user.WalletReply{}
	err = dao.Wallet.Ctx(ctx).Where("uid", req.Uid).Scan(reply)
	return reply, err

}
