package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListPublicWithdraw(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *withdraw.ListPublicWithdrawRes
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListPublicWithdraw(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPublicWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
