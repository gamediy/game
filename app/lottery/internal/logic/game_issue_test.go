package logic

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"reflect"
	"testing"
)

func TestGetIssue(t *testing.T) {
	type args struct {
		ctx      context.Context
		gameCode int32
	}
	tests := []struct {
		name    string
		args    args
		want    GameIssueRespone
		wantErr bool
	}{
		{
			args: args{
				ctx:      context.Background(),
				gameCode: 2000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIssue(tt.args.ctx, tt.args.gameCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIssue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
