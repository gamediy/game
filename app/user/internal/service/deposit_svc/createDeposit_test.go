package deposit_svc

import (
	"context"
	"game/utility/utils/xpusher"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestCreateDeposit_Exec(t *testing.T) {
	type fields struct {
		PayId           int
		Amount          float64
		TransferOrderNo string
		TransferImg     string
		Lang            string
		Uid             int64
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{fields: fields{
			PayId:           3,
			Amount:          10,
			TransferOrderNo: "123",
			TransferImg:     "img",
			Lang:            "en",
			Uid:             161,
		}},
	}
	xpusher.InitFromCfg(gctx.New())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := CreateDeposit{
				PayId:           tt.fields.PayId,
				Amount:          tt.fields.Amount,
				TransferOrderNo: tt.fields.TransferOrderNo,
				TransferImg:     tt.fields.TransferImg,
				Lang:            tt.fields.Lang,
				Uid:             tt.fields.Uid,
			}
			got, err := input.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
