package gamesvc

import (
	"context"
	"game/app/game/api/game/game"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListGameCategory(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *game.ListGameCategoryRes
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListGameCategory(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListGameCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
