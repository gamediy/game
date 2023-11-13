package cq9

import (
	"context"
	"game/app/third/api/cq9/v1"

	cq9svc "game/app/third/internal/service/cq9"
)

func (c *ControllerV1) CheckPlayer(ctx context.Context, req *v1.CheckPlayerReq) (res *v1.CheckPlayerRes, err error) {
	ok := cq9svc.CheckPlayer(ctx, req.Account)
	return &v1.CheckPlayerRes{
		Data:   ok,
		Status: v1.SetStatusOk(),
	}, nil
}
