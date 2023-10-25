package get

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
)

func Role(ctx context.Context, id int64) (*entity.Role, error) {
	var d entity.Role
	_ = dao.Role.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, fmt.Errorf("role not exist")
	}
	return &d, nil
}
