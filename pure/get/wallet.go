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

func Wallet(ctx context.Context, uid int64) (*entity.Wallet, error) {
	var d entity.Wallet
	_ = dao.Wallet.Ctx(ctx).Scan(&d, "uid", uid)
	if d.Uid == 0 {
		return nil, consts.ErrDataNotFound
	}
	return &d, nil
}
func TransactionTypeFromCache(ctx context.Context, code int) (*entity.TransactionType, error) {
	var d entity.TransactionType
	one, err := gcache.GetOrSetFunc(ctx, fmt.Sprint("transType", code), func(ctx context.Context) (value interface{}, err error) {
		return dao.TransactionType.Ctx(ctx).One("code", code)
	}, time.Minute*10)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Scan(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
