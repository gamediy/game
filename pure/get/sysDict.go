package get

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

func Dict(ctx context.Context, key string) (*entity.Dict, error) {
	var d entity.Dict
	_ = dao.Dict.Ctx(ctx).Scan(&d, "k", key)
	if d.Id == 0 {
		return nil, fmt.Errorf("data not found")
	}
	return &d, nil
}
func ImgPrefix() string {
	ctx := context.TODO()
	v, err := gcache.GetOrSetFunc(ctx, "", func(ctx context.Context) (value interface{}, err error) {
		dict, err := Dict(ctx, "cloudflare_pub")
		return dict.V, err
	}, time.Minute*10)
	if err != nil {
		g.Log().Error(ctx, err)
		return ""
	}
	return v.String()
}
