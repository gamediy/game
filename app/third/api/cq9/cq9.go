// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package cq9

import (
	"context"
	
	"game/app/third/api/cq9/v1"
)

type ICq9V1 interface {
	CheckPlayer(ctx context.Context, req *v1.CheckPlayerReq) (res *v1.CheckPlayerRes, err error)
	Balance(ctx context.Context, req *v1.BalanceReq) (res *v1.BalanceRes, err error)
}


