package user

import (
	"context"
	"game/app/user/api/user/wallet"
	"game/app/user/internal/service/wallet_svc"
)

func (*Controller) ListChangeLog(ctx context.Context, req *wallet.ListChangeLogReq) (res *wallet.ListChangeLogRes, err error) {
	x := wallet_svc.ListChangeLog{
		Uid:  req.Uid,
		Page: int(req.Page),
		Size: int(req.Size),
	}
	return x.Exec(ctx)
}
