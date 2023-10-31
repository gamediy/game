package sys

import (
	"game/app/gateway/internal/svc/sys_svc"
	"game/utility/utils/res"
	"github.com/gogf/gf/v2/net/ghttp"
)

func UploadFile(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		token = r.GetQuery("token").String()
	)
	fileName, err := sys_svc.UploadFile(ctx, token)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(fileName, r)
}
