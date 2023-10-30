package withdraw_svc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"testing"
)

func TestBindWithdrawAccount_Exec(t *testing.T) {
	type fields struct {
		BankId  int64
		Address string
		Title   string
		Pass    string
		Uid     int64
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
		{fields: fields{
			BankId:  1,
			Address: "awefalfjdf",
			Title:   "test",
			Pass:    "123456",
			Uid:     161,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BindWithdrawAccount{
				BankId:  tt.fields.BankId,
				Address: tt.fields.Address,
				Title:   tt.fields.Title,
				Pass:    tt.fields.Pass,
				Uid:     tt.fields.Uid,
			}
			if err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
