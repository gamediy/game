package gamesvc

import (
	"context"
	"game/app/game/api/game/game"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListGame(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    *game.ListGameRes
		wantErr bool
	}{
		{args: args{code: "100"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListGame(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
