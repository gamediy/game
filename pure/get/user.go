package get

import (
	"backend/consts"
	"backend/db/dao"
	"backend/db/model/entity"
	"context"
)

func User(ctx context.Context, account string) (*entity.User, error) {
	var d entity.User
	_ = dao.User.Ctx(ctx).Scan(&d, "account", account)
	if d.Id == 0 {
		return nil, consts.ErrAccountNotFound
	}
	return &d, nil
}
