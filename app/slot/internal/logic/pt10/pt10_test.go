package pt10

import (
	"context"
	"game/db/model/entity"
	"reflect"
	"testing"
)

func TestSpin(t *testing.T) {
	type args struct {
		ctx  context.Context
		play entity.Play
		game entity.Game
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
		wantErr bool
	}{
		{
			args: args{
				ctx: context.TODO(),
				play: entity.Play{
					Number: "1001",
				},
				game: entity.Game{
					Number: "1000",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, _, err := Spin(tt.args.ctx, tt.args.play, tt.args.game)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Spin() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
