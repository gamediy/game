package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/pure/get"
)

func GetPayPassStatus(ctx context.Context, in *withdraw.PayPassStatusReq) (*withdraw.PayPassStatusRes, error) {
	u, err := get.UserById(ctx, in.Uid)
	if err != nil {
		return nil, err
	}

	res := &withdraw.PayPassStatusRes{}
	res.Status = "0"
	if u.PayPass != "" {
		res.Status = "1"
	}
	return res, nil
}
