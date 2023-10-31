package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/app/user/api/user/wallet"
)

func (userSvc) Wallet(ctx context.Context, req *user.WalletRequest) (*user.WalletReply, error) {
	return userClient.Wallet(ctx, req)
}
func (userSvc) ListChangeLog(ctx context.Context, req *wallet.ListChangeLogReq) (*wallet.ListChangeLogRes, error) {
	return walletClient.ListChangeLog(ctx, req)
}
