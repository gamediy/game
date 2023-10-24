package user_svc

import (
	"context"
	"game/app/user/api/deposit/deposit"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListDepositAmountItems(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *deposit.DepositAmountItemsRes
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListDepositAmountItems(tt.args.ctx)
			g.Dump(got)
		})
	}
}
