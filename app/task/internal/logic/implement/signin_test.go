package implement

import (
	"context"
	"game/model"
	"testing"
)

func TestSignIn(t *testing.T) {
	type args struct {
		ctx  context.Context
		task model.TaskItem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx: context.TODO(),
				task: model.TaskItem{
					Uid:     121,
					Account: "join",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SignIn(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
