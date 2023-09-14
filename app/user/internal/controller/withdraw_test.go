package controller

import (
	"context"
	withdraw_pb "game/app/user/pb/withdraw"
	"reflect"
	"testing"
)

func TestWithdraw_Submit(t *testing.T) {
	type fields struct {
		UnimplementedWithdrawServiceServer withdraw_pb.UnimplementedWithdrawServiceServer
	}
	type args struct {
		ctx context.Context
		in  *withdraw_pb.SubmitRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantReply *withdraw_pb.SubmitReply
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Withdraw{
				UnimplementedWithdrawServiceServer: tt.fields.UnimplementedWithdrawServiceServer,
			}
			gotReply, err := s.Submit(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReply, tt.wantReply) {
				t.Errorf("Submit() gotReply = %v, want %v", gotReply, tt.wantReply)
			}
		})
	}
}
