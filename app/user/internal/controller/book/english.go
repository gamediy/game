package book

import (
	"context"
	"game/app/user/api/book/english"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (*Controller) GetEnglish(ctx context.Context, req *english.TestEnReq) (res *english.TestEnRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
