package user_svc

import (
	"context"
	"game/app/user/api/user/user"
)

func (userSvc) Wallet(ctx context.Context, req *user.WalletRequest) (*user.WalletReply, error) {
	return userClient.Wallet(ctx, req)
}
