package xtime

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestNow(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NowForISO8601()
			g.Dump(got)
		})
	}
}
