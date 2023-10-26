package get

import (
	"context"
	"fmt"
	"game/consts"
	"game/db/dao"
	"game/db/model/entity"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

func AmountItem(ctx context.Context, id int64) (*entity.AmountItem, error) {
	var d entity.AmountItem
	_ = dao.AmountItem.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, consts.ErrDataNotFound
	}
	return &d, nil
}
func AmountItemFromCache(ctx context.Context, id int64) (*entity.AmountItem, error) {
	var d entity.AmountItem
	data, err := gcache.GetOrSetFunc(ctx, fmt.Sprint("AmountItemFromCache", id), func(ctx context.Context) (value interface{}, err error) {
		return AmountItem(ctx, id)
	}, time.Minute*consts.NormalCacheTime)
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	_ = data.Scan(&d)

	return &d, nil
}
