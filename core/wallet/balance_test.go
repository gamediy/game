package wallet

import (
	"context"
	"github.com/bwmarrin/snowflake"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func TestBalanceUpdate_Update(t *testing.T) {
	node, _ := snowflake.NewNode(1)

	orderNo := node.Generate()
	type fields struct {
		Uid             int64
		Amount          int64
		Profit          int64
		OrderNoRelation int64
		BalanceCode     int
		Title           string
		Note            string
	}
	type args struct {
		ctx context.Context
		fc  func(ctx context.Context, tx gdb.TX) error
	}
	gtime.Now()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{},
			fields: fields{
				Uid:             1,
				BalanceCode:     Freeze,
				OrderNoRelation: orderNo.Int64(),
				Title:           "title",
				Note:            "自己写",
				Amount:          1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Balance{
				Uid:    tt.fields.Uid,
				Amount: float64(tt.fields.Amount),

				OrderNoRelation: tt.fields.OrderNoRelation,
				BalanceCode:     tt.fields.BalanceCode,

				Note: tt.fields.Note,
			}
			if err := this.Update(tt.args.ctx, tt.args.fc); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
