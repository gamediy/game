package mailbox_svc

import (
	"context"
	"game/app/user/api/user/mailbox"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListMailBox(t *testing.T) {
	type args struct {
		ctx context.Context
		req *mailbox.ListMailBoxReq
	}
	tests := []struct {
		name    string
		args    args
		want    *mailbox.ListMailBoxRes
		wantErr bool
	}{
		{args: args{ctx: context.TODO(), req: &mailbox.ListMailBoxReq{Page: 1, Size: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListMailBox(tt.args.ctx, tt.args.req)
			if err != nil {
				t.Fatal(err)
			}
			g.Dump(got)
		})
	}
}
