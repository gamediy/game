package logic

import (
	"context"
	"game/utility/xetcd"
	"testing"
)

func Test(t *testing.T) {

	xetcd.InitEtcd(context.TODO())
	Test123(context.TODO())

}
