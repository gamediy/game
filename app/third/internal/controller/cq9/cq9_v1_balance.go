package cq9

import (
	"context"
	cq9svc "game/app/third/internal/service/cq9"

	"game/app/third/api/cq9/v1"
)

func (c *ControllerV1) Balance(ctx context.Context, req *v1.BalanceReq) (res *v1.BalanceRes, err error) {
	balance, err := cq9svc.Balance(ctx, req.Account)
	if err != nil {
		return nil, err
	}
	return &v1.BalanceRes{
		Data: struct {
			Balance  float64 `json:"balance"`
			Currency string  `json:"currency"`
		}{
			Balance:  balance,
			Currency: "CNY",
		},
		Status: v1.SetStatusOk(),
	}, nil
}
