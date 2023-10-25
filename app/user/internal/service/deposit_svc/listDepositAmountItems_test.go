package deposit_svc

import (
	"context"
	"game/app/user/api/user/deposit"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

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
			got, err := ListDepositAmountItems(tt.args.ctx, 121)
			if err != nil {
				t.Fatal(err)
			}
			g.Dump(got)
		})
	}
}
