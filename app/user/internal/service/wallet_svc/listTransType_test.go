package wallet_svc

import (
	"context"
	"game/app/user/api/user/wallet"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListTransType_Exec(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *wallet.ListTransTypeRes
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := ListTransType{}
			got, err := li.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
