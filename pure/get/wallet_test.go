package get

import (
	"backend/db/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestTransactionTypeFromCache(t *testing.T) {
	type args struct {
		ctx  context.Context
		code int
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.TransactionType
		wantErr bool
	}{
		{args: args{code: 200}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransactionTypeFromCache(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionTypeFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}

func TestWallet(t *testing.T) {
	type args struct {
		ctx context.Context
		uid int64
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Wallet
		wantErr bool
	}{
		{args: args{uid: 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Wallet(tt.args.ctx, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
