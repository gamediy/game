package echo

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"game/app/third/api/echo/v1"
)

func (c *ControllerV1) Say(ctx context.Context, req *v1.SayReq) (res *v1.SayRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
