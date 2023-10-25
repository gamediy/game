package get

import (
	"context"
	"game/consts"
	"game/db/dao"
	"game/db/model/entity"
)

func User(ctx context.Context, account string) (*entity.User, error) {
	var d entity.User
	_ = dao.User.Ctx(ctx).Scan(&d, "account", account)
	if d.Id == 0 {
		return nil, consts.ErrAccountNotFound
	}
	return &d, nil
}

func UserById(ctx context.Context, id int64) (*entity.User, error) {
	var d entity.User
	_ = dao.User.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, consts.ErrAccountNotFound
	}
	return &d, nil

}
