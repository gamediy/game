package cq9

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestCheckPlayer(t *testing.T) {
	type args struct {
		ctx     context.Context
		account string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				account: "test001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPlayer(tt.args.ctx, tt.args.account)
			g.Dump(got)
		})
	}
}
