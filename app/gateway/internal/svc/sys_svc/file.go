package sys_svc

import (
	"context"
	"fmt"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/user/api/user/user"
	"game/utility/utils/xfile"
	"github.com/gogf/gf/v2/frame/g"
)

func UploadFile(ctx context.Context, token string) (g.Map, error) {
	info, err := user_svc.Service.UserInfo(ctx, &user.UserInfoRequest{Token: token})
	if err != nil {
		return nil, err
	}
	if info.Uid == 0 {
		return nil, fmt.Errorf("not auth")
	}
	x := xfile.NewCloudFlareFromCtx(ctx)
	fileName, err := x.Upload(ctx, 5)
	if err != nil {
		return nil, err
	}

	return g.Map{
		"fileName":  fileName,
		"imgPrefix": x.ImgPrefix,
		"url":       x.ImgPrefix + fileName,
	}, nil
}
