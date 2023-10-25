package mailbox_svc

import (
	"context"
	"testing"
)

func TestSentMailBox_Exec(t *testing.T) {
	type fields struct {
		Title    string
		Content  string
		Receiver int64
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{fields: fields{
			Title:    "hello3",
			Content:  "hello3",
			Receiver: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SentMailBox{
				Title:    tt.fields.Title,
				Content:  tt.fields.Content,
				Receiver: tt.fields.Receiver,
			}
			if err := m.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
