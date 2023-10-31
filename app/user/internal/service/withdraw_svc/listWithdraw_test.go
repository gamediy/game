package withdraw_svc

import (
	"context"
	"reflect"
	"testing"
)

func TestListWithdraw_Exec(t *testing.T) {
	type fields struct {
		Uid  int64
		Page int
		Size int
		Lang string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *withdraw.ListWithdrawRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ListWithdraw{
				Uid:  tt.fields.Uid,
				Page: tt.fields.Page,
				Size: tt.fields.Size,
				Lang: tt.fields.Lang,
			}
			got, err := m.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() got = %v, want %v", got, tt.want)
			}
		})
	}
}
