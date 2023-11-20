package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"testing"
)

func TestUpdateLoginPass(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *user.UpdateLoginPassReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				in: &user.UpdateLoginPassReq{
					Uid:     121,
					OldPass: "a1231231",
					NewPass: "123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateLoginPass(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLoginPass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
