package get

import (
	"backend/consts"
	"backend/db/dao"
	"backend/db/model/entity"
	"context"
)

func Admin(ctx context.Context, account string) (*entity.Admin, error) {
	var admin entity.Admin
	_ = dao.Admin.Ctx(ctx).Scan(&admin, "account", account)
	if admin.Id == 0 {
		return nil, consts.ErrAccountNotFound
	}
	return &admin, nil
}
func AdminById(ctx context.Context, id int64) (*entity.Admin, error) {
	var d entity.Admin
	_ = dao.Admin.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, consts.ErrAccountNotFound
	}
	return &d, nil
}
