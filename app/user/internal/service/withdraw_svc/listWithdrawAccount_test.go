package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListWithdrawAccount_Exec(t *testing.T) {
	type fields struct {
		Uid  int64
		Page int
		Size int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *withdraw.ListWithdrawAccountRes
		wantErr bool
	}{
		{
			fields: fields{Uid: 161},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ListWithdrawAccount{
				Uid:  tt.fields.Uid,
				Page: tt.fields.Page,
				Size: tt.fields.Size,
			}
			got, err := m.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
