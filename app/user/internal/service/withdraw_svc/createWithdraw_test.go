package withdraw_svc

import (
	"context"
	"game/utility/utils/xpusher"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestCreateWithdraw_Exec(t *testing.T) {
	type fields struct {
		AmountItemId      int
		Amount            float64
		WithdrawAccountId int
		Lang              string
		Uid               int64
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				Amount:            10,
				AmountItemId:      3,
				WithdrawAccountId: 1727906595624456192,
				Lang:              "",
				Uid:               161,
			},
		},
	}

	xpusher.InitFromCfg(gctx.New())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CreateWithdraw{
				AmountItemId:      tt.fields.AmountItemId,
				Amount:            tt.fields.Amount,
				WithdrawAccountId: tt.fields.WithdrawAccountId,
				Lang:              tt.fields.Lang,
				Uid:               tt.fields.Uid,
			}
			if err := m.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
