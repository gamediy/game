package cq9

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestBalance(t *testing.T) {
	type args struct {
		ctx     context.Context
		account string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{args: args{account: "Tom"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Balance(tt.args.ctx, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
