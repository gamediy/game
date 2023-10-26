package deposit_svc

import (
	"context"
	"game/app/user/api/user/deposit"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListDeposit(t *testing.T) {
	type args struct {
		ctx context.Context
		req *deposit.ListDepositReq
	}
	tests := []struct {
		name    string
		args    args
		want    *deposit.ListDepositRes
		wantErr bool
	}{
		{args: args{
			ctx: context.TODO(),
			req: &deposit.ListDepositReq{
				Uid:  156,
				Size: 2,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListDeposit(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
