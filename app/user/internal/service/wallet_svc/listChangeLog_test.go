package wallet_svc

import (
	"context"
	"game/app/user/api/user/wallet"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListChangeLog_Exec(t *testing.T) {
	type fields struct {
		Uid       int64
		Page      int
		Size      int
		TransCode string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *wallet.ListChangeLogRes
		wantErr bool
	}{
		{
			fields: fields{Uid: 161, Size: 3, TransCode: "-100"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ListChangeLog{
				Uid:       tt.fields.Uid,
				Page:      tt.fields.Page,
				Size:      tt.fields.Size,
				TransCode: tt.fields.TransCode,
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
