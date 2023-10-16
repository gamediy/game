package slot_svc

import (
	"context"
	"game/app/slot/api/slot/slot"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var (
	Service = &slotSvc{}
)

var (
	conn       *grpc.ClientConn
	slotClient slot.SlotServiceClient
)

func SlotClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("slot_svc", grpcx.Balancer.WithRandom())
	slotClient = slot.NewSlotServiceClient(conn)

}
func (slotSvc) Spin(ctx context.Context, req *slot.SpinReq) (res *slot.SpinRes, err error) {

	return slotClient.Spin(ctx, req)

}
func (slotSvc) CheckWon(ctx context.Context, req *slot.CheckWonReq) (res *slot.CheckWonRes, err error) {

	return slotClient.CheckWon(ctx, req)

}

type slotSvc struct {
}
