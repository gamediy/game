package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListWithdrawMethod_Exec(t *testing.T) {
	type fields struct {
		Uid  int64
		Lang string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *withdraw.ListWithdrawMethodRes
		wantErr bool
	}{
		{
			fields: fields{
				Uid:  161,
				Lang: "en",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ListWithdrawMethod{
				Uid:  tt.fields.Uid,
				Lang: tt.fields.Lang,
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
