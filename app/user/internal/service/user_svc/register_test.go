package user_svc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"testing"
)

func TestRegister_Exec(t *testing.T) {
	type fields struct {
		Account     string
		Password    string
		Xid         string
		Phone       string
		Email       string
		Ip          string
		RealName    string
		ClientAgent string
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
				Account: "1234567",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Register{
				Account:     tt.fields.Account,
				Password:    tt.fields.Password,
				Xid:         tt.fields.Xid,
				Phone:       tt.fields.Phone,
				Email:       tt.fields.Email,
				Ip:          tt.fields.Ip,
				RealName:    tt.fields.RealName,
				ClientAgent: tt.fields.ClientAgent,
			}
			if _, err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
